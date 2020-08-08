// length slider
$(function() {
	var rangePercent = $('[id="slider-len"]').val();
	$('[id="slider-len"]').on('change input', function() {
		rangePercent = $('[id="slider-len"]').val();
		$('[id="num-len"]').html(+rangePercent+5);
		$('[id="slider-len"], [id="num-len"]>span').css('filter', 'hue-rotate(-' + rangePercent*100/27 + 'deg)');
	});
});

// extra security slider
$(function() {
	var rangePercent = $('[id="slider-extra"]').val();
	$('[id="slider-extra"]').on('change input', function() {
		rangePercent = $('[id="slider-extra"]').val();
		$('[id="num-extra"]').html(rangePercent);
		$('[id="slider-extra"], [id="num-extra"]>span').css('filter', 'hue-rotate(-' + rangePercent*10 + 'deg)');
	});
});
