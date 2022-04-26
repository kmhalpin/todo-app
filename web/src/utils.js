export const dateFormat = new Intl.DateTimeFormat('id', {
  dateStyle: 'medium', timeStyle: 'medium'
});

export const getDateString = (date) => dateFormat.format(date);