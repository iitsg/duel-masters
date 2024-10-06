package dm07

import (
	"duel-masters/game/civ"
	"duel-masters/game/family"
	"duel-masters/game/fx"
	"duel-masters/game/match"
)

func CosmicNebula(c *match.Card) {

	c.Name = "Cosmic Nebula"
	c.Power = 3000
	c.Civ = civ.Water
	c.Family = []string{family.CyberVirus}
	c.ManaCost = 5
	c.ManaRequirement = []string{civ.Water}

	c.Use(fx.Creature, fx.Evolution, fx.When(fx.MyDrawStep, fx.MayDraw1))

}

func CuriousEye(c *match.Card) {

	c.Name = "Curious Eye"
	c.Power = 1000
	c.Civ = civ.Water
	c.Family = []string{family.CyberVirus}
	c.ManaCost = 3
	c.ManaRequirement = []string{civ.Water}

	c.Use(fx.Creature, fx.WheneverThisAttacksMayLookAtOpShield())

}
