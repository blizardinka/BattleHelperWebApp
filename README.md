##Project Layout
## grouped by Function("Layered Architecture", MVC, etc)
>presentation
>business logic
>external dependencies
data - sample data
handlers - http handlers 
model - for defining our combination and reviews models
##grouped by Module
>cards
>storage
>reviews
Bad idea cuz of naming clash
If we have Json struct and json what ever else - we can't use both
## grouped by Context
For example cards
>Hero cards
    - atack
    - health
    - card ability
>Minion cards
    - ability power
## hexagonal Architecture