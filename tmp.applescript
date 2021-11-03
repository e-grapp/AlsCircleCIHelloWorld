tell application "System Events"
	tell application "Safari" to activate
	delay 5
	tell process "Safari"
		set frontmost to true
		click menu item "Preferences…" of menu 1 of menu bar item "Safari" of menu bar 1
		click button "Advanced" of toolbar 1 of window 1
		tell checkbox "Show Develop menu in menu bar" of group 1 of group 1 of window 1
			if value is 0 then click it
		end tell
		click button 1 of window 1
		click menu item "Allow Remote Automation" of menu 1 of menu bar item "Develop" of menu bar 1
	end tell
end tell