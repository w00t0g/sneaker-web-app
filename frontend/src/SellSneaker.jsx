import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import Dialog from './Dialog'

const SellSneaker = () => {
  const [sneakers, setSneakers] = useState([])
  const [error, setError] = useState('')
  const [selected, setSelected] = useState(null)
  const [formData, setFormData] = useState({
    quantity: 0,
    price: 0
  })
  const [dialogMessage, setDialogMessage] = useState(null);

  const navigate = useNavigate()

  useEffect(() => {
    fetch('http://localhost:8080/api/sneakers')
      .then(res => res.json())
      .then(data => setSneakers(data))
      .catch(error => setError(error))
  }, [])

  const handleSelectChange = (e) => {
    setSelected(e.target.value)
  }

  const handleInputChange = (e) => {
    console.log(formData)
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    })
  }

  const handleSubmit = (e) => {
    formData.price = parseFloat(formData.price)
    formData.quantity = parseInt(formData.quantity)
    e.preventDefault()
    fetch(`http://localhost:8080/api/sneakers/${selected}/sell`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(formData)
    })
      .then(res => res.json())
      .then(data => {
        if (data.error) {
          setError(data.error)
        } else {
          // alert('Sneaker sold')
          // navigate('/sold')
          setDialogMessage('Sneaker sold successfully');
        }
        
      })
      .catch(error => console.log(error))
  }

  const handleDialogClose = () => {
    setDialogMessage(null);
    navigate('/sold');
}

  return (
    <div>
      <h2>Sell Sneaker</h2>
      {error && <p style={{ color: "red" }}>{error}</p>}
      {dialogMessage && (
                <Dialog message={dialogMessage} onClose={handleDialogClose} />
            )}
      <div style={{ display: "flex", alignItems: "center" }}>
        <div>Select Model</div>
        <select style={{ marginLeft: "20px" }} name="select" id="select" onChange={handleSelectChange}>
          <option value="">Select a sneaker</option>
          {sneakers.map(sneaker => (
            <option key={sneaker.id} value={sneaker.id}>{sneaker.model}</option>
          ))}
        </select>
      </div>
      <div>
        {selected && sneakers.filter(sneaker => sneaker.id === parseInt(selected)).map(sneaker => (
          <div key={sneaker.id}>
            <form onSubmit={handleSubmit}>
              <p>Brand: {sneaker.brand}</p>
              <p>Model: {sneaker.model}</p>
              <p>Color: {sneaker.color}</p>
              <p>Platform: {sneaker.platform}</p>
              <p>Available Quantity: {sneaker.quantity}</p>
              <label htmlFor="quantity">Quantity</label>
              <input onChange={handleInputChange} type="number" id="quantity" name="quantity" min="1" max={sneaker.quantity} />
              <label htmlFor="price">Price</label>
              <input onChange={handleInputChange} type="number" id="price" name="price" min="1" />
              <button>Sell</button>
            </form>
          </div>

        ))}
      </div>
    </div>
  )
}

export default SellSneaker