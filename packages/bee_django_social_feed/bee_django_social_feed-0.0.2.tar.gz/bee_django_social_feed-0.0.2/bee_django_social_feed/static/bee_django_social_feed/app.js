$(document).ready(function () {
    var new_feed = new Vue({
        el: '#new-feed',
        data: {
            content: '',
            success_message: '',
            error_message: ''
        }
    });

    var feeds = new Vue({
        el: '#feeds',
        data: {
            feeds: [],
            page: 1
        }
    });

    var loader = new Vue({
        el: '#feeds-loader',
        data: {
            status: 0
        },

        methods: {
            load_more_page: function (e) {
                e.preventDefault();

                feeds.page++;
                loader.status = 1;
                load_page_data();
            }
        }
    });

    Vue.component('feed', {
        template: '#feed-template',
        props: ['feed']
    });

    moment.locale('zh-cn');
    Vue.filter('formatDate', function (value) {
        if (value) {
            return moment(String(value)).fromNow();
        }
    });


    $('#post-new-feed').click(function (e) {
        e.preventDefault();
        $(this).prop('disabled', true);
        var url = $(this).attr('data-confirm');

        if (new_feed.content === '') {
            alert('内容不能为空');
        } else {
            data = {content: new_feed.content};
            $.post(url, data).done(function (data) {
                new_feed.success_message = data['message'];

                new_feed.content = '';
                setTimeout(function () {
                    new_feed.success_message = '';
                }, 2000);

                new_feeds = JSON.parse(data['new_feeds']);
                feeds.feeds = new_feeds.concat(feeds.feeds);
            })
        }
        $(this).prop('disabled', false);
    });

    // 拉取指定页面的feeds
    var load_page_data = function () {
        $.get('/feed/feeds', {page: feeds.page})
            .done(function (data) {
                new_feeds = JSON.parse(data['feeds']);
                if (new_feeds.length === 0) {
                    feeds.page = data['page'];
                    loader.status = 3;
                } else {
                    feeds.feeds = feeds.feeds.concat(new_feeds);
                    feeds.page = data['page'];
                    loader.status = 0;
                }

                // console.log('第' + feeds.page + '页拉取完毕');
            })
            .fail(function () {
                loader.status = 2;
            })
    };
    load_page_data();

    // 下拉翻页的处理
    window.onscroll = function (e) {
        var bottomOfWindow = document.documentElement.scrollTop
            + window.innerHeight === document.documentElement.offsetHeight;

        if (bottomOfWindow) {
            loader.load_more_page(e);
        }
    }
});


