import React from "react";

// keep this as separate to cuisines since both will need to be different shortly
class Descriptors extends React.Component {
    // Constructor
    constructor(props) {
        super(props);

        this.state = {
            descriptors: [],
            DataisLoaded: false,
        };
    }

    // ComponentDidMount is used to
    // execute the code
    componentDidMount() {
        fetch("http://localhost:8080/descriptors")
            .then((res) => res.json())
            .then((json) => {
                this.setState({
                    descriptors: json,
                    DataisLoaded: true,
                });
            });
    }

    render() {
        const { DataisLoaded, descriptors } = this.state;
        if (!DataisLoaded)
            return (
                <div>
                    <h1> Pleses wait some time.... </h1>
                </div>
            );

        return (
            <div className="container">
                <h2>Descriptors</h2>
                {descriptors.map((descriptor) => (
                    <div>
                        <ol>
                            <div>{descriptor.Name} ({descriptor.Count})</div>
                        </ol>
                    </div>
                ))}
            </div>
        );
    }
}

export default Descriptors;
