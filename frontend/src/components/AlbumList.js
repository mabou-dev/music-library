import React, { useEffect } from "react"
import { useDispatch, useSelector } from "react-redux";
import { fetchAlbums } from "../redux/albumSlice";

function AlbumList() {
    const dispatch = useDispatch();
    const { list: albums, status } = useSelector((state) => state.albums);

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
