// app.js
import React, { useState, useEffect } from 'react';
import MovieList from './MovieList.js';
import NewMovieForm from './NewMovieForm';

const App = () => {
  // State to store the list of movies
  const [movies, setMovies] = useState([]);

  // Fetch the list of movies from the backend on component mount
  useEffect(() => {
    fetchMovies();
  }, []);

  // Function to fetch the list of movies from the backend
  const fetchMovies = async () => {
    try {
      const response = await fetch('https://sturdy-doodle-vrgrv655qvp379v-3000.app.github.dev/movies');
      const data = await response.json();
      setMovies(data);
    } catch (error) {
      console.error('Error fetching movies:', error);
    }
  };

  // Function to add a new movie to the backend
  const addMovie = async (newMovie) => {
    try {
      const response = await fetch('https://sturdy-doodle-vrgrv655qvp379v-3000.app.github.dev/movies', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(newMovie),
      });

      if (response.ok) {
        // If the movie is added successfully, fetch the updated list
        fetchMovies();
      } else {
        console.error('Failed to add movie');
      }
    } catch (error) {
      console.error('Error adding movie:', error);
    }
  };

  return (
    <div>
      <h1>Movie Watchlist</h1>
      {/* Component to display the list of movies */}
      <MovieList movies={movies} />
      {/* Component to add new movies */}
      <NewMovieForm onAddMovie={addMovie} />
    </div>
  );
};

export default App;
