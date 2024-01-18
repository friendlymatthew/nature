## Visual Realism
This program displays an open leaf venation model using the space colonization algorithm. Since I don't leave my room, might as well draw some nature.
[The Space Colonization Paper](http://algorithmicbotany.org/papers/colonization.egwnp2007.large.pdf)

### Pictures
<img width="497" alt="Screen Shot 2024-01-18 at 4 53 00 PM" src="https://github.com/friendlymatthew/nature/assets/38759997/e746a9b0-0b28-4b10-ba9c-a08f82276932">


### Libraries used
Uses [Ebitengine](https://github.com/hajimehoshi/ebiten) and Go's [math](https://pkg.go.dev/math) library. 
Vector computation is done in the `vec` package.

### More work
- [ ] create a spatial indexer and query via knn to find the closest node quickly
- [ ] **thicken** the veins [Murray's Law](https://en.wikipedia.org/wiki/Murray%27s_law)
- [ ] allow color + # of nodes, attractors, dist options
- [ ] write a closed venation model
- [ ] transpose audio into attractors, model music with venation
- [ ] **go outside** or up the vitamin d dosage
