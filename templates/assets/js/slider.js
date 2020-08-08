// length slider
$(function() {
	var rangeVal = $('[id="slider-len"]').val();

	$('[id="slider-len"]').on('change input', function() {
		rangeVal = parseInt($('[id="slider-len"]').val());
		$('[id="len"]').html(5 + parseInt($('[id="slider-len"]').val()) + parseInt($('[id="slider-extra"]').val())); // update len
		$('[id="slider-len"], [id="len"]>span').css('filter', 'hue-rotate(-' + rangeVal*100/27 + 'deg)');
	});
});

// extra security slider
$(function() {
	var rangeVal = $('[id="slider-extra"]').val();

	$('[id="slider-extra"]').on('change input', function() {
		rangeVal = parseInt($('[id="slider-extra"]').val());
		$('[id="extra"]').html(rangeVal);

        if (document.getElementById('slider-len').disabled === true) {
            $('[id="len"]').html(parseInt($('[id="slider-len"]').val()) + parseInt($('[id="slider-extra"]').val())); // update len
        } else {
            $('[id="len"]').html(5 + parseInt($('[id="slider-len"]').val()) + parseInt($('[id="slider-extra"]').val())); // update len
        }

		$('[id="slider-extra"], [id="extra"]>span').css('filter', 'hue-rotate(-' + rangeVal*10 + 'deg)');
	});
});
