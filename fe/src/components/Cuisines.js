import React from "react";

class Cuisines extends React.Component {
    // Constructor
    constructor(props) {
        super(props);

        this.state = {
            cuisines: [],
            DataisLoaded: false,
        };
    }

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
        const { DataisLoaded, cuisines } = this.state;
        if (!DataisLoaded)
            return (
                <div>
                    <h1> Pleses wait some time.... </h1>
                </div>
            );

        return (
            <div className="container">
                <h2>Cuisines</h2>
                {cuisines.map((cuisine) => (
                    <div>
                        <ol>
                            <div>{cuisine.Name} ({cuisine.Recipes})</div>
                        </ol>
                    </div>
                ))}
            </div>
        );
    }
}

export default Cuisines;
