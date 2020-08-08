$(function() {
	var rangePercent = $('[type="range"]').val();
	$('[type="range"]').on('change input', function() {
		rangePercent = $('[type="range"]').val();
		$('h4').html(+rangePercent+5+'<span></span>');
		$('[type="range"], h4>span').css('filter', 'hue-rotate(-' + rangePercent*100/27 + 'deg)');
		$('h4').css({'transform': 'translateX(-50%) scale(' + (1.25+(rangePercent/75)) + ')', 'left': rangePercent*100/27+'%'});
	});
});
