package routes

// User Timelines:
// GET /api/tweets/user/:userId - Get tweets posted by a specific user.
// GET /api/tweets/following - Get tweets from users the logged-in user is following.
// GET /api/tweets/public - Get public tweets from all users.
// GET /api/tweets/trending - Get trending tweets based on popular hashtags.

// Timeline Pagination:
// GET /api/tweets/user/:userId?page=:pageNumber&limit=:pageSize - Get paginated tweets for a specific user.
// GET /api/tweets/following?page=:pageNumber&limit=:pageSize - Get paginated tweets from users the logged-in user is following.
// GET /api/tweets/public?page=:pageNumber&limit=:pageSize - Get paginated public tweets
