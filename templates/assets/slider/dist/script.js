$(function() {
	var rangePercent = $('[type="range"]').val();
	$('[type="range"]').on('change input', function() {
		rangePercent = $('[type="range"]').val();
		$('h4').text(rangePercent);
		$('[type="range"]').css('filter', 'hue-rotate(-' + rangePercent + 'deg)');
	});
});