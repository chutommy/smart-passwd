The Scrolling Background:

	This relies entirely on CSS to do its thing, which is cool, but that makes
	changing it a bit weird/tricky at first. You can still use pretty much any image
	you want, but for best results make sure yours is:

	- Horizontally tileable.
	- Wide and short.
	- About 1500px wide.
	- Fades to a solid color either at the top of bottom (which is used to fill
	  the empty space above or below your image).

	Now, there are two ways to use it: with CSS, or with Sass:

	CSS:

		Look for this line in css/style.css (line 108 as of this writing):

			background: #348cb2 url("images/bg.jpg") bottom left;

		and use it to set the page background color, URL, and placement of
		your image. It should be as close to 1500px wide as you can get it.

	Sass:

		Set the value of $bg to the page background color, URL, and placement
		of your image. Change $bg-width if your image is something other than
		1500px wide.
