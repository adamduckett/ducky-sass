# Ducky Sass
This is a sample Sass project architecture using the [Atomic design](http://bradfrost.com/blog/post/atomic-web-design/) methodology.

At the moment, apart from [`normalize.css`](http://necolas.github.io/normalize.css/) resets, there are no styles included with Ducky Sass; only a directory structure. However, this may change in the future if I feel including example Sass would be beneficial.

Each folder contains its own `README.md` file explaining its purpose.

What Ducky Sass does include, though, are basic `package.json` and `gulpfile.js` files which after running `npm install` in our directory will get us fully
setup with a fully automated watch task which watches for changes and compiles
our Sass.

## Structure

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
