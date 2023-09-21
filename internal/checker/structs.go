package checker


type InviteRes struct {
    Type                        int         `json:"type"`
    Code                        string      `json:"code"`
    Guild                       Guild       `json:"guild"`
    GuildID                     string      `json:"guild_id"`
    ApproximateMemberCount      int         `json:"approximate_member_count"`
    ApproximatePresenceCount    int         `json:"approximate_presence_count"`
}

type Guild struct {
    Name                     string   `json:"name"`
    VerificationLevel        int      `json:"verification_level"`
    PremiumSubscriptionCount int      `json:"premium_subscription_count"`
}