

/*
Function sends AJAX GET request every 1000ms to update total 
accumulated time for timer specified
*/
function updateStatus() {
	$.ajax({
		url: "/status",
		data: {timerName: $("#myTimer").val()},
		dataType: "json",
		type: "GET",
		success: function (data, startTime) {
			$(".holder").html("Total accumulated time in seconds: " + data.totaltime);
		}
	});
	setTimeout(updateStatus, interval);
}


/*
Function sends AJAX POST request using start time, timer name 
and a boolean indicating if the timer is new
*/
function startTimer() {
	$("#myButton").html("Stop!");
	$(".timerName").html("Timer name: " + $("#myTimer").val());
	var startTime = new Date;
	var timerId = setInterval(function() {
		$(".timerDisplay").text("Current active timer: " + (new Date - startTime) / 1000);
	}, 50);
	document.getElementById("myButton").onclick = function() { stopTimer(timerId); };
	$.ajax({
		url: "/start",
		data: JSON.stringify({timerName: $("#myTimer").val(), startTime: startTime, isNew: false}),
		dataType: "json",
		type: "POST"
	}).done(function(data) {
	});
}


/*
Function sends AJAX POST request using stop time and timer name, then displays accumulated 
time for the timer name specified and clears running timer back to 0.000
*/
function stopTimer(timerId) {
	var stopTime = new Date;
	$.ajax({
		url: "/stop",
		data: JSON.stringify({timername: $("#myTimer").val(), stopTime: stopTime}),
		dataType: "json",
		type: "POST"
	}).success(function(data) {
		clearInterval(timerId);
		$(".timerDisplay").html("Current active timer: 0.000");
		document.getElementById("myButton").onclick = startTimer;
		$("#myButton").html("Start!");
	});
}

setTimeout(updateStatus, 1000);
var interval = 1000;