# Requirements Specification

Hi and welcome to the team. As you know, we are a small shop with a prime location in a
prominent city ran by a friendly shopkeep named Allison. We buy and sell many of the finest goods.
Unfortunately, our goods are constantly decreasing in price as they approach their sell by date. We
have a system in place that updates our inventory for us. It was developed by a no-nonsense type named
Leeroy, who has moved on to new adventures.

Your task is to add a new feature to our system so that
we can begin selling a new category of items, Flowers. First an introduction to our system:

- All items have a `sell_by` value which denotes the number of days we have left to sell the item before it expires
- All items have a `price` value which denotes how much the item costs
- At the end of each day our system adjusts both values for every item

Pretty simple, right? Well this is where it gets interesting:

- Once the sell by date has passed, `price` decreases twice as fast
- The `price` of an item is never negative
- **Fine Art** actually increases in `price` the older it gets
- **Gold Coins**, being an hard asset, never have to be sold and they never decrease in `price`
- The `price` of an item is never more than 50. Except for **Gold Coins**. They are so valuable that their `price` is 80
- **Concert Tickets** increase in `price` the closer `sell_by` gets to zero:
  - `price` increases by 2 when there are 10 days or less and by 3 when there are 5 days or less
  - `price` drops to 0 after the concert

We have recently started selling **flowers**. We need you to update the system and implement the following behavior:

- **Flowers** decrease in `price` twice as fast as normal items

### Interview Guidelines:
* Please ask questions.
* Feel free to Google things during the interview.
* Add/change code as much as you like, as long as everything still works correctly.
* If you do things differently than you would for production code please tell us that.
