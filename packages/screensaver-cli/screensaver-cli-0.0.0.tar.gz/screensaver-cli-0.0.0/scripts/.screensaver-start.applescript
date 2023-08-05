#!/usr/bin/osascript

-- Arabesque
-- Flurry
-- Shell
-- iLifeslideshows
-- Word of the Day
-- iTunes Artwork
-- Computer Name

on run argv
    tell application "System Events"
        -- get name of screen savers
        start (screen saver (item 1 of argv))
    end tell
end run
