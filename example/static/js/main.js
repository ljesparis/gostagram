$(document).ready(function() {
	$("#instagram-user").hover(function() {
		$(this).popover('show')
	}, function() {
		$(this).popover('hide')
	})
})