package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"newsletter/x/newsletter/types"
)

// SetNewsletter set a specific newsletter in the store from its index
func (k Keeper) SetNewsletter(ctx sdk.Context, newsletter types.Newsletter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewsletterKeyPrefix))
	b := k.cdc.MustMarshal(&newsletter)
	store.Set(types.NewsletterKey(
		newsletter.Index,
	), b)
}

// GetNewsletter returns a newsletter from its index
func (k Keeper) GetNewsletter(
	ctx sdk.Context,
	index string,

) (val types.Newsletter, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewsletterKeyPrefix))

	b := store.Get(types.NewsletterKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNewsletter removes a newsletter from the store
func (k Keeper) RemoveNewsletter(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewsletterKeyPrefix))
	store.Delete(types.NewsletterKey(
		index,
	))
}

// GetAllNewsletter returns all newsletter
func (k Keeper) GetAllNewsletter(ctx sdk.Context) (list []types.Newsletter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewsletterKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Newsletter
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
