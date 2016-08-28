import { MajordomoWebPage } from './app.po';

describe('majordomo-web App', function() {
  let page: MajordomoWebPage;

  beforeEach(() => {
    page = new MajordomoWebPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
