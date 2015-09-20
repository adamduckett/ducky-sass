# Ducky Sass
This is a sample Sass project architecture using the [Atomic design](http://bradfrost.com/blog/post/atomic-web-design/) methodology.

Apart from [`normalize.css`](http://necolas.github.io/normalize.css/) resets, Ducky Sass contains no styles; only a project directory structure.

It does, however, come bundled with a `gulpfile.js` file which containing Gulp tasks to both compile our Sass and watch for changes.

I've not included things like `index.html` or `img/` (which I appreciate are common project files and directories) because I just want to focus on the Sass/CSS side of things with Ducky Sass.

## Getting Started

Installing Ducky Sass via the command line is as simple as:

    $ git clone https://github.com/adamduckett/ducky-sass.git your-project-folder
    $ cd your-project-folder
    $ ./start

Running the `start` file removes the `.git` folder before installing the dependencies needed for our Gulp tasks (to compile our Sass) and finally deleting itself. Self-destruction, yo!


## Structure

The `sass/` directory's structure is based on Brad Frost's [Atomic design](http://bradfrost.com/blog/post/atomic-web-design/) methodology.

    sass/
        atoms/
        molecules/
        organisms/
        pages/
        shame/
        utils/
        vendor/
        main.scss

### Molecules
The `molecules/` folder contains styles for more tangible components constructed by combining multiple atoms. Molecules are the backbone of our design system and could include things like a search form built from a label, input and button.

### Organisms
The `organisms/` folder contains relatively complex, distinct sections of a website. This is where our system becomes more visually relatable to clients; think: headers, navigation, sidebars and footers.

### Pages
The `pages/` folder holds page-specific styles. If a rule only affects an element on a certain page then it should go here. Files for what I regard the most common pages are included: 'home', 'about' and 'contact'.

### Shame
The `shame/` folder is optional but I regularly make use of it in my projects. Need to apply a quick fix you totally plan on re-writing later? Put it here. It's especially useful on projects involving multiple developers as it makes the code we're all embarrassed to own up to 'public' for all to see.

### Atoms
The `atoms/` folder contains any high-level unclassed HTML element styles for our project. This is where you define your basic typography rules as well as generic button, list and input styles.

### Vendor
The `vendor/` folder contains external vendor styles your project relies on like Nicolas Gallagher's [`normalize.css`](http://necolas.github.io/normalize.css/).


## License

Copyright 2015 Adam Duckett

Licensed under the Apache License, Version 2.0.
