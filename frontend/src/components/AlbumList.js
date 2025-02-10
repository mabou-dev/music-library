import React, { useEffect } from "react"

const albums = [
    {
        id: "a",
        title: "title_A",
        artist: "artist_A"
    },
    {
        id: "b",
        title: "title_B",
        artist: "artist_B"
    }
]

function AlbumList() {
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
