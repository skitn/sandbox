class Product
  @name = "Product"

  def self.name
    @name
  end

  def initialize(name)
    @name = name
  end

  def name
    @name
  end
end

class DVD < Product
  @name = "DVD"

  def self.name
    @name
  end

  def upcase_name
    @name.upcase
  end
end

p Product.name

product = Product.new("Instance Product")
p product.name

p Product.name

p DVD.name
dvd = DVD.new("Instance DVD")
p dvd.name
p dvd.upcase_name

p DVD.name
