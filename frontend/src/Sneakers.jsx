import React, { useState, useEffect } from 'react'
import Dialog from './Dialog'

const Sneakers = () => {
  const [sneakers, setSneakers] = useState([])
  const [error, setError] = useState('')
  const [dialogMessage, setDialogMessage] = useState(null);

  useEffect(() => {
    fetch('http://localhost:8080/api/sneakers')
      .then(response => response.json())
      .then(data => setSneakers(data))
      .catch(error => setError(error))

  }, [])
  const deleteSneaker = (id) => {
    return () => {
      fetch(`http://localhost:8080/api/sneakers/${id}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json'
        }
      })
        .then(response => response.json())
        .then(data => {
          if (data.error) {
            setError(data.error)
          } else {
            // alert('Sneaker deleted successfully')
            setDialogMessage('Sneaker deleted successfully');
            setSneakers(sneakers.filter(sneaker => sneaker.id !== id))
          }
        })
        .catch(error => setError(error))
    }
  }

  const handleDialogClose = () => {
    setDialogMessage(null);
}

  return (
    <div>
      <h2>Sneaker Inventory</h2>
      {error && <p>{error}</p>}
      {dialogMessage && (
                <Dialog message={dialogMessage} onClose={handleDialogClose} />
            )}
      <small>
      <table>
        <thead>
          <tr>
            <th>Purchase Date</th>
            <th>Brand</th>
            <th>Model</th>
            <th>Color</th>
            <th>Platform</th>
            <th>Quantity</th>
            <th>Purchase Price</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          {sneakers.length === 0 && <tr><td colSpan="8">No sneakers in inventory</td></tr>}
          {sneakers.map(sneaker => (
            <tr key={sneaker.id}>
              <td>{sneaker.purchaseDate}</td>
              <td>{sneaker.brand}</td>
              <td>{sneaker.model}</td>
              <td>{sneaker.color}</td>
              <td>{sneaker.platform}</td>
              {/* <td>{sneaker.quantity}</td> */}
              {/* if quantity is zero it is sold out */}
              <td>{sneaker.quantity === 0 ? (<span style={{backgroundColor:"red", padding:"0 3px"}}>Sold out</span>) : sneaker.quantity}</td>
              <td>{sneaker.purchasePrice}</td>
              <th><button onClick={deleteSneaker(sneaker.id)} style={{backgroundColor:"red"}}>Delete</button></th>
            </tr>
          ))}
        </tbody>
      </table>
      </small>    
    </div>
  )
}

export default Sneakers