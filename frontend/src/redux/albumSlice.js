import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";

export const fetchAlbums = createAsyncThunk("albums/fetchAlbums", async () => {
    const response = await fetch("/api/albums");
    return response.json();
});

const albumSlice = createSlice({
    name: "albums",
    initialState: {
        list: [],
        status: "idle",
    },
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(fetchAlbums.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchAlbums.fulfilled, (state, action) => {
                state.list = action.payload;
                state.status = "succeeded";
            })
            .addCase(fetchAlbums.rejected, (state) => {
                state.status = "failed";
            });
    },
});

export default albumSlice.reducer;
