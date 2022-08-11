class Inventory

  def initialize(items)
    @items = items
  end

  def update_price
    @items.each do |item|
      item.adjust_price
      if item.name != "Fine Art" and item.name != "Concert Tickets"
        if item.price > 0
          if item.name != "Gold Coins"
            item.price = item.price - 1
          end
          if item.name == 'Flowers' && item.price > 0
            item.price = item.price - 1
          end
        end
      else
        if item.price < 50
          item.price = item.price + 1
          if item.name == "Concert Tickets"
            if item.sell_by < 11
              if item.price < 50
                item.price = item.price + 1
              end
            end
            if item.sell_by < 6
              if item.price < 50
                item.price = item.price + 1
              end
            end
          end
        end
      end
      if item.name != "Gold Coins"
        item.sell_by = item.sell_by - 1
      end
      if item.sell_by < 0
        if item.name != "Fine Art"
          if item.name != "Concert Tickets"
            if item.price > 0
              if item.name != "Gold Coins"
                item.price = item.price - 1
              end
              if item.name == 'Flowers' && item.price > 0
                item.price = item.price - 1
              end
            end
          else
            item.price = item.price - item.price
          end
        else
          if item.price < 50
            item.price = item.price + 1
          end
        end
      end
    end
  end
end

class Gold < Item

  def adjust_price
  end
end

class Flower < Item
  PRICE_ADJUSTMENT = -2
  PRICE_ADJUSTMENT_MODIFIER = 2

  def adjust_price
    if sell_by > 0
      price = price - PRICE_ADJUSTMENT
    else
      price = price - (PRICE_ADJUSTMENT * PRICE_ADJUSTMENT_MODIFIER)
    end
    price = 0 if price < 0
    adjust_sell_by
  end
end

GenericItem

class Item
  attr_accessor :name, :sell_by, :price

  def initialize(name, sell_by, price)
    @name = name
    @sell_by = sell_by
    @price = price
  end

  def adjust_price
    if price > 0
      price = price - 1
    end
  end

  def adjust_sell_by
    sell_by = sell_by - 1
  end

  def to_s
    "#{@name}, #{@sell_by}, #{@price}"
  end
end
