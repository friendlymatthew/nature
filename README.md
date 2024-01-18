## Visual Realism

### 
[The Space Colonization Paper](http://algorithmicbotany.org/papers/colonization.egwnp2007.large.pdf)

This program displays an open leaf venation model using the space colonization algorithm. Since I don't leave my room, might as well draw some nature.

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