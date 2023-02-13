package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"newsletter/x/newsletter/types"
)

// SetNewsletterInfo set newsletterInfo in the store
func (k Keeper) SetNewsletterInfo(ctx sdk.Context, newsletterInfo types.NewsletterInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewsletterInfoKey))
	b := k.cdc.MustMarshal(&newsletterInfo)
	store.Set([]byte{0}, b)
}

// GetNewsletterInfo returns newsletterInfo
func (k Keeper) GetNewsletterInfo(ctx sdk.Context) (val types.NewsletterInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewsletterInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNewsletterInfo removes newsletterInfo from the store
func (k Keeper) RemoveNewsletterInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewsletterInfoKey))
	store.Delete([]byte{0})
}
