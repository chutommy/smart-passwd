// length slider
$(function () {
  let $lenSlider = jQuery('[id="slider-len"]');
  let rangeVal = $lenSlider.val();
  $('[id="slider-len"]').on('change input', function () {
    rangeVal = parseInt($lenSlider.val());
    $('[id="len"]').html(5 + parseInt($lenSlider.val()) + parseInt($('[id="slider-extra"]').val())); // update len
    if (rangeVal < 20) {
      $('[id="slider-len"], [id="len"]>span').css('filter', 'hue-rotate(' + rangeVal * 100 / 27 * 1.5 + 'deg)');
    } else {
      $('[id="slider-len"], [id="len"]>span').css('filter', 'hue-rotate(' + 100 + 'deg)');
    }
  });
});

// extra security value slider
$(function () {
  let $extraSlider = jQuery('[id="slider-extra"]');
  let rangeVal = $extraSlider.val();
  $extraSlider.on('change input', function () {
    rangeVal = parseInt($extraSlider.val());
    $('[id="extra"]').html(rangeVal);
    if (document.getElementById('slider-len').disabled === true) {
      $('[id="len"]').html(parseInt($('[id="slider-len"]').val()) + parseInt($('[id="slider-extra"]').val())); // update len
    } else {
      $('[id="len"]').html(5 + parseInt($('[id="slider-len"]').val()) + parseInt($('[id="slider-extra"]').val())); // update len
    }
    if (rangeVal < 6) {
      $('[id="slider-extra"], [id="extra"]>span').css('filter', 'hue-rotate(' + rangeVal * 10 * 2 + 'deg)');
    } else {
      $('[id="slider-extra"], [id="extra"]>span').css('filter', 'hue-rotate(' + 100 + 'deg)');
    }
  });
});
