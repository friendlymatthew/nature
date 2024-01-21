## Visual Realism

This program displays an open leaf venation model using the space colonization algorithm. Since I don't leave my room, might as well procedurally generate trees.

![tree](https://github.com/friendlymatthew/nature/assets/38759997/7acd5af8-6967-41f9-be9d-8798d78af8b7)


### Libraries used + References
Uses [Ebitengine](https://github.com/hajimehoshi/ebiten) and Go's [math](https://pkg.go.dev/math) library. 
[The Space Colonization Paper](http://algorithmicbotany.org/papers/colonization.egwnp2007.large.pdf)

Vector computation is done in the `vec` package.

### Todos
- [x] host a demo
- [ ] optimize finding the closest node
> This would be very cool. A database of attractors using [quad tree](https://www.geeksforgeeks.org/quad-tree/#) and querying using KNN.
- [ ] **thicken** the veins [Murray's Law](https://en.wikipedia.org/wiki/Murray%27s_law)
- [ ] allow color + # of nodes, attractors, dist options
- [ ] write a closed venation model
- [ ] transpose audio into attractors, model music with venation
- [ ] **go outside** or up the vitamin d dosage
