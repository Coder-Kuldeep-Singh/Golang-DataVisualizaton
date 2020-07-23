// taking data from server and then parsing according to need
var GetALLData = async () => {
  const response = await fetch("/api/json/data");
  const myJson = await response.json(); //extract JSON from the http response
  //   console.log(myJson.Visual);
  var cloud = "";
  for (var idx = 0; idx < myJson.Visual.length; idx++) {
    cloud += myJson.Visual[idx].topic + " ";
  }
  //   console.log(cloud);
  // ZC.LICENSE = ["569d52cefae586f634c54f86dc99e6a9", "b55b025e438fa8a98e32482b5f768ff5"];
  let chartConfig = {
    type: "wordcloud",
    options: {
      text: cloud,
      aspect: "spiral",
      colorType: "palette",
      ignore: ["establish", "this"],
      maxItems: 50,
      minLength: "4px",
      palette: [
        "#D32F2F",
        "#1976D2",
        "#9E9E9E",
        "#E53935",
        "#1E88E5",
        "#7E57C2",
        "#F44336",
        "#2196F3",
        "#A1887F",
      ],
      rotate: true,
      style: {
        tooltip: {
          text: "Frequncy %hits",
          padding: "5px",
          alpha: 0.9,
          backgroundColor: "#D32F2F",
          borderColor: "none",
          borderRadius: "3px",
          fontColor: "white",
          fontFamily: "Georgia",
          textAlpha: 1,
          visible: true,
          width: "100px",
          wrapText: true,
        },
        fontFamily: "Merriweather",
        hoverState: {
          backgroundColor: "#1976D2",
          borderColor: "none",
          borderRadius: "3px",
          fontColor: "white",
        },
      },
    },
  };

  zingchart.render({
    id: "myChart",
    data: chartConfig,
  });
};

// calling function to render word cloud chart
GetALLData();
