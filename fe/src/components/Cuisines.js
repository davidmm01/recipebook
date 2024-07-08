import React from "react";
import "./cuisines.css";

// used the following to help make the button group
// https://dev.to/ramonak/react-how-to-create-a-custom-button-group-component-in-5-minutes-3lfd

class Cuisines extends React.Component {
  // Constructor
  constructor(props) {
    super(props);

    this.state = {
      cuisines: [],
      DataisLoaded: false,
      selectedCuisineKey: -1,
    };
  }

  setSelectedCuisineKey = (key) => {
    console.log("eat my eggs, key=", key);
    // on re-click of already selected cuisine, clear selection
    if (this.state.selectedCuisineKey === key) {
      key = -1;
    }
    this.setState({ selectedCuisineKey: key });
  };

  // ComponentDidMount is used to
  // execute the code
  componentDidMount() {
    fetch("http://localhost:8080/cuisines")
      .then((res) => res.json())
      .then((json) => {
        this.setState({
          cuisines: json,
          DataisLoaded: true,
        });
      });
  }

  render() {
    const { DataisLoaded, cuisines, selectedCuisineKey } = this.state;
    if (!DataisLoaded)
      return (
        <div>
          <h1> Pleses wait some time.... </h1>
        </div>
      );

    return (
      <div className="container">
        <h2>Cuisines</h2>
        {cuisines.map((cuisine, i) => (
          <div>
            <ol>
              <button
                key={i}
                name={cuisine.Name}
                onClick={() => this.setSelectedCuisineKey(i)}
                className={
                  i === this.state.selectedCuisineKey
                    ? "customButton active"
                    : "customButton"
                }
              >
                {cuisine.Name} ({cuisine.Count})
              </button>
            </ol>
          </div>
        ))}
      </div>
    );
  }
}

export default Cuisines;
