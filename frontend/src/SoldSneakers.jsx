import React, { useState, useEffect } from 'react'

const SoldSneakers = () => {
    const [soldSneakers, setSoldSneakers] = useState([])
    const [error, setError] = useState('')

    useEffect(() => {
        fetch('http://localhost:8080/api/sold-sneakers')
            .then(res => res.json())
            .then(data => setSoldSneakers(data))
            .catch(error => setError(error))
    }, [])

    return (
        <div>
            <h2>Sold Sneakers</h2>
            <div style={{ color: "red" }}>
                {/* {error && <p>{error}</p>} */}
            </div>
            <small>
                <table>
                    <thead>
                        <tr>
                            <th>Brand</th>
                            <th>Model</th>
                            <th>Color</th>
                            <th>Platform</th>
                            <th>Quantity</th>
                            <th>Purchase Price</th>
                            <th>Sale Price</th>
                            <th>Result</th>
                        </tr>
                    </thead>
                    <tbody>
                        {soldSneakers.length === 0 ? (<tr><td colSpan="6">No sneakers sold</td></tr>) : (
                            <>
                                {soldSneakers.map(soldSneaker => (
                                    <tr key={soldSneaker.id}>
                                        <td>{soldSneaker.sneaker.brand}</td>
                                        <td>{soldSneaker.sneaker.model}</td>
                                        <td>{soldSneaker.sneaker.color}</td>
                                        <td>{soldSneaker.sneaker.platform}</td>
                                        <td>{soldSneaker.quantity}</td>
                                        <td>{soldSneaker.sneaker.purchasePrice * soldSneaker.quantity} @{soldSneaker.sneaker.purchasePrice}</td>
                                        <td>{soldSneaker.price * soldSneaker.quantity} @{soldSneaker.price}</td>
                                        <td style={{ color: (soldSneaker.price - soldSneaker.sneaker.purchasePrice) * soldSneaker.quantity < 0 ? "red" : "green" }}>{(soldSneaker.price - soldSneaker.sneaker.purchasePrice) * soldSneaker.quantity}</td>
                                    </tr>
                                ))}
                                <tr style={{ fontWeight: "bold" }}>
                                    <td colSpan="7">Total</td>
                                    {/* <td>{soldSneakers.reduce((acc, soldSneaker) => acc + soldSneaker.sneaker.purchasePrice * soldSneaker.quantity, 0)}</td> */}
                                    {/* <td>{soldSneakers.reduce((acc, soldSneaker) => acc + soldSneaker.price * soldSneaker.quantity, 0)}</td> */}
                                    <td style={{ color: soldSneakers.reduce((acc, soldSneaker) => acc + (soldSneaker.price - soldSneaker.sneaker.purchasePrice) * soldSneaker.quantity, 0) < 0 ? "red" : "green" }}>
                                        {soldSneakers.reduce((acc, soldSneaker) => acc + (soldSneaker.price - soldSneaker.sneaker.purchasePrice) * soldSneaker.quantity, 0)}
                                    </td>
                                </tr>
                            </>
                        )}
                    </tbody>
                </table>
            </small>
        </div>
    )
}

export default SoldSneakers