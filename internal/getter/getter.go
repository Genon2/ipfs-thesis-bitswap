package getter

import (
	"context"
	"errors"
	"fmt"

	notifications "github.com/Genon2/ipfs-thesis-bitswap/internal/notifications"
	logging "github.com/ipfs/go-log"

	blocks "github.com/ipfs/go-block-format"
	cid "github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
)

var log = logging.Logger("bitswap")

// GetBlocksFunc is any function that can take an array of CIDs and return a
// channel of incoming blocks.
type GetBlocksFunc func(context.Context, []cid.Cid) (<-chan blocks.Block, error)

// SyncGetBlock takes a block cid and an async function for getting several
// blocks that returns a channel, and uses that function to return the
// block syncronously.
func SyncGetBlock(p context.Context, k cid.Cid, gb GetBlocksFunc) (blocks.Block, error) {
	if !k.Defined() {
		log.Error("undefined cid in GetBlock")
		return nil, blockstore.ErrNotFound
	}
	fmt.Printf("[%s] BITSWAP GETTER from SyncGetBlock in getter.go\n", k.String())
	// Any async work initiated by this function must end when this function
	// returns. To ensure this, derive a new context. Note that it is okay to
	// listen on parent in this scope, but NOT okay to pass |parent| to
	// functions called by this one. Otherwise those functions won't return
	// when this context's cancel func is executed. This is difficult to
	// enforce. May this comment keep you safe.
	ctx, cancel := context.WithCancel(p)
	defer cancel()

	promise, err := gb(ctx, []cid.Cid{k})
	if err != nil {
		return nil, err
	}
	select {
	case block, ok := <-promise:
		if !ok {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
				return nil, errors.New("promise channel was closed")
			}
		}
		fmt.Printf("[%s] Promise from SyncGetBlock\n", block)
		return block, nil
	case <-p.Done():
		return nil, p.Err()
	}
}

// WantFunc is any function that can express a want for set of blocks.
type WantFunc func(context.Context, []cid.Cid)

// AsyncGetBlocks take a set of block cids, a pubsub channel for incoming
// blocks, a want function, and a close function, and returns a channel of
// incoming blocks.
func AsyncGetBlocks(ctx context.Context, sessctx context.Context, keys []cid.Cid, notif notifications.PubSub,
	want WantFunc, cwants func([]cid.Cid)) (<-chan blocks.Block, error) {

	// If there are no keys supplied, just return a closed channel
	if len(keys) == 0 {
		out := make(chan blocks.Block)
		close(out)
		return out, nil
	}

	// Use a PubSub notifier to listen for incoming blocks for each key
	// Cr??e une liste
	remaining := cid.NewSet()
	promise := notif.Subscribe(ctx, keys...)
	for _, k := range keys {
		log.Debugw("Bitswap.GetBlockRequest.Start", "cid", k)
		remaining.Add(k)
	}
	fmt.Println("Pass through AsynGetBlocks from getter.go")
	// Send the want request for the keys to the network
	// Avertit les Node bitswap
	want(ctx, keys)

	// Cr??e un channel de sortie et on re??oit les valeurs depuis handleIncoming
	out := make(chan blocks.Block)
	go handleIncoming(ctx, sessctx, remaining, promise, out, cwants)
	return out, nil
}

// Listens for incoming blocks, passing them to the out channel.
// If the context is cancelled or the incoming channel closes, calls cfun with
// any keys corresponding to blocks that were never received.
func handleIncoming(ctx context.Context, sessctx context.Context, remaining *cid.Set,
	in <-chan blocks.Block, out chan blocks.Block, cfun func([]cid.Cid)) {

	ctx, cancel := context.WithCancel(ctx)
	// ??a marche pas fmt.Println("[%s] FOUND VIA BITSWAP 42\n", blk.Cid().String())

	// Clean up before exiting this function, and call the cancel function on
	// any remaining keys
	defer func() {
		cancel()
		close(out)
		// can't just defer this call on its own, arguments are resolved *when* the defer is created
		// Cancel les CID jamais re??u par Bitswap et applique la fonction de internal/session/session.go
		cfun(remaining.Keys())
	}()

	for {
		select {
		// blk -> re??oit les donn??es via le stream "in"
		case blk, ok := <-in:
			// If the channel is closed, we're done (note that PubSub closes
			// the channel once all the keys have been received)
			if !ok {
				return
			}
			fmt.Printf("[%s] BLK from handleIncoming in getter.go\n", blk)
			fmt.Printf("[%s] FOUND VIA BITSWAP from handleIncoming in getter.go\n", blk.Cid().String())

			// Re??u via Bitswap le CID
			// retire les CID de remaining
			remaining.Remove(blk.Cid())
			select {
			case out <- blk:
			case <-ctx.Done():
				return
			case <-sessctx.Done():
				return
			}
		case <-ctx.Done():
			return
		case <-sessctx.Done():
			return
		}
	}
}
