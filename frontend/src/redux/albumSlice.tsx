import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';

interface Album {
    id: string;
    title: string;
    artist: string;
}

interface AlbumState {
    albums: Album[];
    status: "idle" | "loading" | "failed";
}

const initialState: AlbumState = {
    albums: [],
    status: "idle",
};

export const fetchAlbums = createAsyncThunk<Album[]>("albums/fetchAlbums", async () => {
    const response = await fetch("/api/albums");
    return response.json();
});

const albumSlice = createSlice({
    name: "albums",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(fetchAlbums.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchAlbums.fulfilled, (state, action: PayloadAction<Album[]>) => {
                state.albums = action.payload;
                state.status = "idle";
            })
            .addCase(fetchAlbums.rejected, (state) => {
                state.status = "failed";
            });
    },
});

export default albumSlice.reducer;
