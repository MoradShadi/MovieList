// NewMovieForm.js
import React, { useState } from 'react';

const NewMovieForm = ({ onAddMovie }) => {
  const [newMovieTitle, setNewMovieTitle] = useState('');

  const handleInputChange = (event) => {
    setNewMovieTitle(event.target.value);
  };

  const handleSubmit = (event) => {
    event.preventDefault();

    if (newMovieTitle.trim() === '') {
      return; // Prevent submitting empty movie titles
    }

    const newMovie = {
      title: newMovieTitle,
    };

    // Call the onAddMovie function passed from the parent component
    onAddMovie(newMovie);

    // Clear the input field
    setNewMovieTitle('');
  };

  return (
    <div>
      <h2>Add a New Movie</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Movie Title:
          <input
            type="text"
            value={newMovieTitle}
            onChange={handleInputChange}
          />
        </label>
        <button type="submit">Add Movie</button>
      </form>
    </div>
  );
};

export default NewMovieForm;
