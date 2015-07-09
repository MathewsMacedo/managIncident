$(document).ready(function () {

// recupBadge();
  //Flash notice
  $(".alert").slideDown(100).fadeIn(1000).removeClass("hide");

  window.setTimeout(function() {
      $(".alert").fadeTo(1000, 0).slideUp(500, function(){
          $(this).remove();
      });
  }, 4000);

  $('.glyphicon-calendar').click(function(){
    if ($('.dtp1 > input').val() != ''){
      $(".dtp2 > input").prop('disabled', false);
    };
  });


  $('.table').dataTable();

  moment.locale('fr');
  $('.datetimepicker').datetimepicker({
    format :"DD-MM-YYYY HH:mm:ss ZZ",
    showTodayButton : true,
    // minDate : new Date(),
  });

  $('.dateRemove2').click(function(){
    $('.inputRemove2').val('01-01-0001 00:00:00 +0000');
  });

  $('.dateRemove1').click(function(){
    $('.inputRemove1').val('01-01-0001 00:00:00 +0000');
  });

  $('.addUser').click(function(){
    	$("#myModal-Add").modal('show');


  });

    $('.pwd').keyup(function(){
      var regex1 = $(this).val();
      var passwordRegex1 = /^[a-z0-9_-]{6,18}$/;
      if ( regex1.match(passwordRegex1)){
        $('.pwd').css({
          'background-color' : '#5CB85C',
          color : '#fff'
        });
      }else{
        $('.pwd').css({
          'background-color' : '#D9534F',
          color : '#fff'
        });
      }
    });

    var valid = 0;
    $('.repwd').keyup(function(){
      var regex2 = $(this).val();
      var passwordRegex2 = /^[a-z0-9_-]{6,18}$/;
      var pass = $('.pwd').val();
      var repass = $('.repwd').val();
      if ((pass == repass) && regex2.match(passwordRegex2) ) {
        $('.repwd').css({
          'background-color' : '#5CB85C',
          color : '#fff'
        });
        valid = 1;
      }else{
        $('.repwd').css({
          'background-color' : '#D9534F',
          color : '#fff'
        });
        valid = 0;
      }
    });

    $('.formPassLogin').submit(function(e){

      if (valid == 0){
        return false;
      }else if (valid == 1){
        return true;
      }
    })

});
