// Dialog.js
import React, { useRef, useEffect } from 'react';

const Dialog = ({ message, onClose }) => {
    const dialogRef = useRef(null);

    useEffect(() => {
        if (dialogRef.current) {
            dialogRef.current.showModal();
        }
    }, []);

    const handleClose = () => {
        if (dialogRef.current) {
            dialogRef.current.close();
        }
        onClose();
    };

    return (
        <dialog ref={dialogRef}>
            <p>{message}</p>
            <menu>
                <button onClick={handleClose}>OK</button>
            </menu>
        </dialog>
    );
};

export default Dialog;
