// var GetALLData = async () => {
//   const response = await fetch("/api/json/data");
//   const myJson = await response.json(); //extract JSON from the http response
//   // console.log(myJson.Visual);
//   return myJson;
// };

function GetLocation(event) {
  var location = event.target.value;
  console.log(location);
  var perpage = document.getElementById("perpage");
  var values = parseInt(perpage.value);

  var body = document.getElementById("table_body").innerHTML;
  if (location != "") {
    var New_location = location.toLowerCase();

    var GetALLData = async () => {
      const response = await fetch("/api/json/data");
      const myJson = await response.json(); //extract JSON from the http response
      // console.log(myJson.Visual);
      // var cities = [];

      // for (var idx = 0; idx < myJson.Visual.length; idx++) {
      //   // console.log(myJson.Visual[idx].title);
      //   cities.push(myJson.Visual[idx].title);
      // }
      // console.log(cities);

      var dropdown_locations = [
        {
          Title: "",
          Topic: "",
          Start_Year: 0,
          Intensity: 0,
          Sector: "",
          Region: "",
          Pestle: "",
        },
      ];

      for (var i = 0; i < myJson.Visual.length; i++) {
        var loc = myJson.Visual[i].title.toLowerCase();
        var exists = loc.includes(New_location); //if value is anywhere is the string
        //   // var exists = loc.startsWith(New_location); // if the string start with given string
        if (exists) {
          dropdown_locations.push({
            Title: myJson.Visual[i].title,
            Topic: myJson.Visual[i].topic,
            Start_Year: myJson.Visual[i].start_year,
            Intensity: myJson.Visual[i].intensity,
            Sector: myJson.Visual[i].sector,
            Region: myJson.Visual[i].region,
            Pestle: myJson.Visual[i].pestle,
          });
        }
      }
      // console.log(dropdown_locations.);

      var dropdown_length = dropdown_locations.length;
      // console.log(dropdown_length);

      // // Remove all existing elements
      // var existing = document.querySelector(".drop-down");
      var body = document.getElementById("table_body");
      // existing = document.querySelector(".drop-down");
      // body.innerHTML = "";
      var table = "";
      // var para = document.createElement("datalist");
      // para.id = "drop";
      if (dropdown_length > 25) {
        for (var j = 1; j <= values; j++) {
          // console.log(dropdown_locations.Title);
          if (location != "") {
            table += `<tr>
                    <td style="padding: 15px;">${dropdown_locations[j].Title}</td>
                    <td style="padding: 15px;">${dropdown_locations[j].Topic}</td>
                    <td style="padding: 15px;">${dropdown_locations[j].Start_Year}</td>
                    <td style="padding: 15px;">${dropdown_locations[j].Intensity}</td>
                    <td style="padding: 15px;">${dropdown_locations[j].Sector}</td>
                    <td style="padding: 15px;">${dropdown_locations[j].Region}</td>
                    <td style="padding: 15px;">${dropdown_locations[j].Pestle}</td>
                  </tr>`;
          } else {
            body.innerHTML = body;
            return;
          }
        }
      } else {
        if (location != "") {
          for (var j = 1; j <= dropdown_length; j++) {
            // Show(drop, New_location, existing, para);
            table += `<tr>
                      <td style="padding: 15px;">${dropdown_locations[j].Title}</td>
                      <td style="padding: 15px;">${dropdown_locations[j].Topic}</td>
                      <td style="padding: 15px;">${dropdown_locations[j].Start_Year}</td>
                      <td style="padding: 15px;">${dropdown_locations[j].Intensity}</td>
                      <td style="padding: 15px;">${dropdown_locations[j].Sector}</td>
                      <td style="padding: 15px;">${dropdown_locations[j].Region}</td>
                      <td style="padding: 15px;">${dropdown_locations[j].Pestle}</td>
                    </tr>`;
          }
        } else {
          body.innerHTML = body;
          return;
        }
      }
      body.innerHTML = table;
    };
    GetALLData();
  } else {
    console.log(body);
    body.innerHTML = body;
  }
}

// function Show(drop, location, existing, para) {
//   if (location != "") {
//     var submenu = document.createElement("option");
//     submenu.value = drop;
//     para.appendChild(submenu);
//     existing.appendChild(para);
//   } else {
//     existing.innerHTML = "";
//   }
// }

function Default_Table() {
  var GetALLData = async () => {
    const response = await fetch("/api/json/data");
    const myJson = await response.json(); //extract JSON from the http response
    var remains = document.getElementById("remaining");
    remains.innerHTML = `<span class="badge badge-dark" style="padding:10px;">
      Showing 1 to 10 of 1,000 entries
      </span>`;
    // console.log('Showing 1 to 10 of 1,000 entries');
    // console.log(structured_data)
    var body = document.getElementById("table_body");
    var table = "";
    for (var i = 1; i <= 10; i++) {
      // console.log(structured_data[i].Title);
      table += `<tr>
          <td style="padding: 15px;">${myJson.Visual[i].title}</td>
          <td style="padding: 15px;">${myJson.Visual[i].topic}</td>
          <td style="padding: 15px;">${myJson.Visual[i].start_year}</td>
          <td style="padding: 15px;">${myJson.Visual[i].intensity}</td>
          <td style="padding: 15px;">${myJson.Visual[i].sector}</td>
          <td style="padding: 15px;">${myJson.Visual[i].region}</td>
          <td style="padding: 15px;">${myJson.Visual[i].pestle}</td>
        </tr>`;
    }
    body.innerHTML = table;
  };
  GetALLData();
}

function SelectFilter(options) {
  var selectedValue = options.value;

  var GetALLData = async () => {
    const response = await fetch("/api/json/data");
    const myJson = await response.json(); //extract JSON from the http response
    var remains = document.getElementById("remaining");
    remains.innerHTML = `<span class="badge badge-dark"  style="padding:10px;">
      Showing 1 to ${selectedValue} of ${myJson.Visual.length} entries
    </span>`;
    // console.log(`Showing 1 to ${selectedValue} of ${structured_data.length - 1} entries`);
    // console.log(structured_data)
    var body = document.getElementById("table_body");
    var table = "";
    console.log("selected value", selectedValue);
    for (var i = 0; i <= selectedValue; i++) {
      table += `<tr>
                <td style="padding: 15px;">${myJson.Visual[i].title}</td>
                <td style="padding: 15px;">${myJson.Visual[i].topic}</td>
                <td style="padding: 15px;">${myJson.Visual[i].start_year}</td>
                <td style="padding: 15px;">${myJson.Visual[i].intensity}</td>
                <td style="padding: 15px;">${myJson.Visual[i].sector}</td>
                <td style="padding: 15px;">${myJson.Visual[i].region}</td>
                <td style="padding: 15px;">${myJson.Visual[i].pestle}</td>
        </tr>`;
    }
    body.innerHTML = table;
  };
  GetALLData();
}
