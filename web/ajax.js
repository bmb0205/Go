

setTimeout(updateStatus, 1000);
interval = 1000;
function updateStatus() {
	$.ajax({
		url: "/status",
		data: {timerName: $("#myTimer").val()},
		dataType: "json",
		type: "GET",
		success: function (data) {
			$("#holder").html("Total accumulated time in seconds: " + data.totaltime);
			var runningTime = (new Date - startTime) / 1000
			$(".timerDisplay").text("Current active timer: " + runningTime);
		}
	});
	setTimeout(updateStatus, interval);
}



function startTimer() {
	$("#myButton").html("Stop!");
	$("#timerName").html("Timer name: " + $("#myTimer").val());
	var startTime = new Date;
	var timerId = setInterval(function() {
		$(".timerDisplay").text("Current active timer: " + (new Date - startTime) / 1000);
	}, 50);
	// document.getElementById("myButton").onclick = function() { stopTimer(timerId); };
	$("#myButton").click(function() { stopTimer(timerId); });
	$.ajax({
		url: "/start",
		data: JSON.stringify({timerName: $("#myTimer").val(), startTime: startTime, isNew: false}),
		dataType: "json",
		type: "POST"
	}).done(function(data) {
		// alert("POST request start endpoint hit");
	});
}



function stopTimer(timerId) {
	var stopTime = new Date;
	$.ajax({
		url: "/stop",
		data: JSON.stringify({timername: $("#myTimer").val(), stopTime: stopTime}),
		dataType: "json",
		type: "POST"
	}).done(function(data) {
		// alert("POST request stop endpoint hit");
	});
	clearInterval(timerId);
	$(".timerDisplay").text("Current active timer: 0.000");
	$("#myButton").click(startTimer);
	// document.getElementById("myButton").onclick = startTimer;
	$("#myButton").html("Start!");
}
