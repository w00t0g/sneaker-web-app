import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import Dialog from './Dialog'

const AddSneaker = () => {
    const [formData, setFormData] = useState({
        brand: '',
        model: '',
        platform: '',
        color: '',
        size: '',
        purchasePrice: '',
        purchaseDate: '',
        quantity: ''
    })
    const [error, setError] = useState('')
    const [dialogMessage, setDialogMessage] = useState(null);

    const navigate = useNavigate()

    const handleChange = (e) => {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value
        })
    }

    const handleSubmit = (e) => {
        e.preventDefault()
        formData.purchasePrice = parseFloat(formData.purchasePrice)
        formData.quantity = parseInt(formData.quantity)
        console.log(formData)

        fetch('http://localhost:8080/api/sneakers', {
            method: 'POST',
            body: JSON.stringify(formData)
        })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                setError(data.error)
            } else {
                // alert('Sneaker added')
                // navigate('/')
                setDialogMessage('Sneaker added successfully');
            }
        })
        .catch(error => {
            setError('Error adding sneaker')
        })
    }

    const handleDialogClose = () => {
        setDialogMessage(null);
        navigate('/');
    }

    return (
        <div>
            <h2>Add Sneaker</h2>
            {error && <p style={{color:"red"}}>{error}</p>}
            {dialogMessage && (
                <Dialog message={dialogMessage} onClose={handleDialogClose} />
            )}
            <form onSubmit={handleSubmit}>
                <label htmlFor="brand">Brand</label>
                <input onChange={handleChange} type="text" name="brand" id="brand" required/>

                <label htmlFor="model">Model</label>
                <input onChange={handleChange} type="text" name="model" id="model" required/>

                {/* platform */}
                <label htmlFor="platform">Platform</label>
                <input onChange={handleChange} type="text" name="platform" id="platform" />

                <label htmlFor="color">Color</label>
                <input onChange={handleChange} type="text" name="color" id="color" required/>

                <label htmlFor="size">Size</label>
                <input onChange={handleChange} type="text" name="size" id="size" required/>

                <label htmlFor="purchasePrice">Purchase Price</label>
                <input onChange={handleChange} type="text" name="purchasePrice" id="purchasePrice"
                    pattern="\d+" title="xx.xx" required
                 />
                 {/* \d+(\.\d{2})? */}

                <label htmlFor="purchaseDate">Purchase Date</label>
                {/* dd/mm/yyyy	 */}
                <input onChange={handleChange} type="text" name="purchaseDate" id="purchaseDate"
                    pattern="\d{1,2}/\d{1,2}/\d{4}" title="dd/mm/yyyy" required
                />

                <label htmlFor="quantity">Quantity</label>
                <input onChange={handleChange} type="text" name="quantity" id="quantity"
                    pattern="\d+" title="Whole number" required
                />

                <button type="submit">Add Sneaker</button>
            </form>
        </div>
    )
}

export default AddSneaker