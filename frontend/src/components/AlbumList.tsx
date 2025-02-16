import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { fetchAlbums } from '../redux/albumSlice';
import { RootState, AppDispatch } from '../redux/store';

function AlbumList() {
    const dispatch = useDispatch<AppDispatch>();
    const albums = useSelector((state: RootState) => state.albums.albums);
    const status = useSelector((state: RootState) => state.albums.status);

    useEffect(() => {
        dispatch(fetchAlbums());
    }, [dispatch]);

    if (status === "loading") return (
        <p>Loading...</p>
    )
    if (status === "failed") return (
        <p>Error fetching albums.</p>
    )

    return (
        <div>
            <h2>Album List</h2>
            <ul>
                {albums.map((album) => (
                    <li key={album.id}>{album.title} - {album.artist}</li>
                ))}
            </ul>
        </div>
    )
}

export default AlbumList;
