use std::fmt;

#[derive(Debug, PartialEq, Eq)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
       let total_minutes = hours * 60 + minutes;
        let normalized_minutes = total_minutes.rem_euclid(24 * 60);
        
        Self {
            hours: (normalized_minutes / 60) % 24,
            minutes: normalized_minutes % 60,
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock::new(self.hours, self.minutes + minutes)
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result{
        write!(f, "{:02}:{:02}", self.hours, self.minutes)
    }
}