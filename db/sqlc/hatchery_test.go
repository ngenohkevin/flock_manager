package db

import (
	"context"
	"github.com/ngenohkevin/flock_manager/db/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomHatchery(t *testing.T, production Production) Hatchery {
	arg := CreateHatcheryParams{
		ProductionID: production.ID,
		Infertile:    util.RandomHatchery(),
		Early:        util.RandomHatchery(),
		Middle:       util.RandomHatchery(),
		Late:         util.RandomHatchery(),
		DeadChicks:   util.RandomHatchery(),
		AliveChicks:  util.RandomHatchery(),
	}
	hatchery, err := testQueries.CreateHatchery(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, hatchery)

	require.Equal(t, arg.ProductionID, hatchery.ProductionID)
	require.Equal(t, arg.Infertile, hatchery.Infertile)
	require.Equal(t, arg.Early, hatchery.Early)
	require.Equal(t, arg.Middle, hatchery.Middle)
	require.Equal(t, arg.Late, hatchery.Late)
	require.Equal(t, arg.DeadChicks, hatchery.DeadChicks)
	require.Equal(t, arg.AliveChicks, hatchery.AliveChicks)

	require.NotZero(t, hatchery.ID)
	require.NotZero(t, hatchery.CreatedAt)

	return hatchery
}

func TestCreateHatchery(t *testing.T) {
	breed := createdRandomBreed(t)
	production := createRandomProduction(t, breed)

	createRandomHatchery(t, production)

}

func TestGetHatchery(t *testing.T) {
	breed := createdRandomBreed(t)
	production := createRandomProduction(t, breed)

	hatchery1 := createRandomHatchery(t, production)

	hatchery2, err := testQueries.GetHatchery(context.Background(), hatchery1.ProductionID)
	require.NoError(t, err)
	require.NotEmpty(t, hatchery2)

	require.Equal(t, hatchery1.ProductionID, hatchery2.ProductionID)
	require.Equal(t, hatchery1.Infertile, hatchery2.Infertile)
	require.Equal(t, hatchery1.Early, hatchery2.Early)
	require.Equal(t, hatchery1.Middle, hatchery2.Middle)
	require.Equal(t, hatchery1.Late, hatchery2.Late)
	require.Equal(t, hatchery1.DeadChicks, hatchery2.DeadChicks)
	require.Equal(t, hatchery1.AliveChicks, hatchery2.AliveChicks)

	require.WithinDuration(t, hatchery1.CreatedAt, hatchery2.CreatedAt, time.Second)
}
