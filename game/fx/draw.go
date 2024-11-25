package fx

import (
	"duel-masters/game/match"
	"fmt"
)

// Draw1 Convenience method with standard signature for drawing 1
func Draw1(card *match.Card, ctx *match.Context) {
	card.Player.DrawCards(1)
}

// Draw2 Convenience method with standard signature for drawing 1
func Draw2(card *match.Card, ctx *match.Context) {
	card.Player.DrawCards(2)
}

// Draw1ToMana draws 1 card and puts it in the players manazone
func Draw1ToMana(card *match.Card, ctx *match.Context) {

	cards := card.Player.PeekDeck(1)

	if len(cards) < 1 {
		return
	}

	c, err := card.Player.MoveCard(cards[0].ID, match.DECK, match.MANAZONE, card.ID)

	if err != nil {
		return
	}

	ctx.Match.ReportActionInChat(card.Player, fmt.Sprintf("%s was added to %s's manazone from the top of their deck", c.Name, card.Player.Username()))

}

func Draw2ToMana(card *match.Card, ctx *match.Context) {
	Draw1ToMana(card, ctx)
	Draw1ToMana(card, ctx)
}

func MayDraw1(card *match.Card, ctx *match.Context) {
	MayDrawAmount(card, ctx, 1)
}

func DrawUpTo2(card *match.Card, ctx *match.Context) {
	DrawBetween(card, ctx, 0, 2)
}

func DrawUpTo3(card *match.Card, ctx *match.Context) {
	DrawBetween(card, ctx, 0, 3)
}

// This gives the player the choice to select a number of cards to draw between 2 provided limits
func DrawBetween(card *match.Card, ctx *match.Context, min int, max int) {
	count := max
	if min != max {
		count = SelectCount(
			card.Player,
			ctx.Match,
			fmt.Sprintf("%s effect: Draw between %d and %d cards", card.Name, min, max),
			min,
			max)
	}
	card.Player.DrawCards(count)
}

// This lets the player choose if they want to draw the full amount or none
func MayDrawAmount(card *match.Card, ctx *match.Context, amount int) {
	drawAmount := 0
	textAmount := fmt.Sprintf("%d cards", amount)
	if amount == 1 {
		textAmount = "1 card"
	}

	if BinaryQuestion(card.Player, ctx.Match, fmt.Sprintf("Do you want to draw %s? (%s effect)", textAmount, card.Name)) {
		drawAmount = amount
	}

	card.Player.DrawCards(drawAmount)
}

// TopCardToShield puts top 1 card from deck to shielzone
func TopCardToShield(card *match.Card, ctx *match.Context) {

	cards := card.Player.PeekDeck(1)

	if len(cards) < 1 {
		return
	}

	_, err := card.Player.MoveCard(cards[0].ID, match.DECK, match.SHIELDZONE, card.ID)

	if err != nil {
		return
	}

	ctx.Match.ReportActionInChat(card.Player, fmt.Sprintf("%s put the top card of his deck into the shieldzone from %s's effect", card.Player.Username(), card.Name))

}
