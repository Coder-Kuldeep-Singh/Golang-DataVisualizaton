var GetALLData = async () => {
  const response = await fetch("/api/json/data");
  const myJson = await response.json(); //extract JSON from the http response
  //   console.log(myJson.Visual);
  var topics = [];
  var intensity = [];
  for (var idx = 0; idx < myJson.Visual.length; idx++) {
    topics.push(myJson.Visual[idx].topic);
    intensity.push(myJson.Visual[idx].intensity);
  }

  const count = (names) =>
    names.reduce((a, b) => ({ ...a, [b]: (a[b] || 0) + 1 }), {}); // don't forget to initialize the accumulator

  let re = count(topics);
  var radar_value = [];
  var radar_labels = [];
  for (var key in re) {
    if (re.hasOwnProperty(key)) {
      // console.log(key); //
      radar_labels.push(key);
      radar_value.push(re[key]);
      //   console.log(re[key]); //
    }
  }
  //   console.log(radar_value);
  //   console.log(radar_labels);

  var marksCanvas = document.getElementById("Radar");

  var marksData = {
    labels: radar_labels,
    datasets: [
      {
        label: "Intensity",
        backgroundColor: "blue",
        data: intensity,
      },
    ],
  };

  var radarChart = new Chart(marksCanvas, {
    type: "radar",
    data: marksData,
  });
};

// calling function to render word cloud chart
GetALLData();
