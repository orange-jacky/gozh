<template>
  <mu-col width="100" tablet="70" desktop="70">
    <mu-paper class="article" :zDepth="2">
      <mu-card>
        <mu-card-title title="扩展包推荐 —— Laravel Log 增强" subTitle="翻译"/>
        <div class="my-article-info">
          <span>Acorns . </span><span>于一周前 . </span><span>8888阅读 . </span><span>原文地址</span>
        </div>
        <div class="article-content">
          <div class="markdown-body topic-content-big-font" id="emojify">

            <blockquote style="margin-top: 0px!important;">
              这是一篇社区协同翻译的文章，已完成翻译，更多信息请点击 <a href="https://laravel-china.org/topics/7440/community-translation-function-online-star2">协同翻译介绍</a> 。
            </blockquote>
            <hr>


            <div class="translation-item active " style="display: block;" index="1" id="SectionIndex_1" name="SectionIndex_1">
              <p>Laravel Log Enhancer 是 <a href="https://laravel-news.com/laravel-5-6">Laravel 5.6</a>&nbsp; 的一个扩展包，可以在 Laravel 日志中添加额外的数据。 得益于 Laravel 5.6 中日志的更新，这个包利用这些特性扩展日志记录并添加数据，例如请求头、内存使用情况、会话信息和其它数据 。</p>

              <div class="postion">


                <div class="text-center ">
                  <a href="https://laravel-china.org/users/2190">
                    <img class="img-thumbnail avatar emoji" src="https://lccdn.phphub.org/uploads/avatars/2190_1439374577.png?imageView2/1/w/100/h/100" style="width:52px; height:52px;margin:5px;border: 1px solid #ddd;">
                  </a>
                  <div class="media-body padding-top-sm">
            <span class="introduction">
                  sork 翻译于 <span class="popover-with-html" data-content="2018-04-05 10:20:20" data-original-title="" title="">1周前</span>
            </span>
                  </div>
                </div>

                <div class="translation-operation-wrap">
                  <div class="index text-center">
                    <a class="section-vote btn btn-sm btn-shadow btn-default  popover-with-html" data-content="多个翻译的情况下，赞数会影响排序" data-url="https://laravel-china.org/translations/sections/2374/vote" data-original-title="" title="">
                      <i class="fa fa-thumbs-up" aria-hidden="true"></i> <span class="vote-count">0</span>
                    </a>

                    <a class="no-pjax btn btn-default btn-shadow btn-sm  popover-with-html" data-content="发现错别字？点此可帮助译者改进" href="https://laravel-china.org/sections/2374/patches/create" data-original-title="" title="">
                      <i class="fa fa-send" aria-hidden="true"></i> 改进
                    </a>

                    <a class="no-pjax btn btn-default btn-shadow btn-sm  no-pjax popover-with-html" data-content="提交你的译文" href="https://laravel-china.org/translations/sections/2374" data-original-title="" title="">
                      <i class="fa fa-paint-brush" aria-hidden="true"></i> 重译
                    </a>

                  </div>

                  <div class="text-center" style="font-size:0.9em;margin-top:12px;color:#a5a5a5;">
                    <i class="fa fa-eye" aria-hidden="true"></i> 由 <a href="https://laravel-china.org/users/1" style="color:#a5a5a5;">Summer</a> 审阅
                  </div>
                </div>
              </div>



              <div class="clearfix"></div>
            </div>
            <div class="translation-item active " style="display: block;" index="2" id="SectionIndex_2" name="SectionIndex_2">
              <p>你可以通过修改 <code>config/logging.php</code> 中的 <code>tap</code> 属性，来添加 &nbsp;<code>LogEnhancer</code>&nbsp;到你的日志通道：</p>
              <pre class="  language-php"><code class="  language-php"><span class="token string">'production_stack'</span> <span class="token operator">=</span><span class="token operator">&gt;</span> <span class="token punctuation">[</span>
    <span class="token string">'driver'</span> <span class="token operator">=</span><span class="token operator">&gt;</span> <span class="token string">'stack'</span><span class="token punctuation">,</span>
    <span class="token string">'tap'</span> <span class="token operator">=</span><span class="token operator">&gt;</span> <span class="token punctuation">[</span><span class="token scope">Freshbitsweb<span class="token punctuation">\</span>LaravelLogEnhancer<span class="token punctuation">\</span>LogEnhancer<span class="token punctuation">::</span></span><span class="token keyword">class</span><span class="token punctuation">]</span><span class="token punctuation">,</span>
    <span class="token string">'channels'</span> <span class="token operator">=</span><span class="token operator">&gt;</span> <span class="token punctuation">[</span><span class="token string">'daily'</span><span class="token punctuation">,</span> <span class="token string">'slack'</span><span class="token punctuation">]</span><span class="token punctuation">,</span>
<span class="token punctuation">]</span><span class="token punctuation">,</span>
</code></pre>
              <p>配置日志通道后，下面是一条包含默认额外信息的日志示例：</p>
              <pre class="  language-php"><code class="  language-php"><span class="token punctuation">[</span><span class="token number">2018</span><span class="token operator">-</span><span class="token number">03</span><span class="token operator">-</span><span class="token number">30</span> <span class="token number">05</span><span class="token punctuation">:</span><span class="token number">07</span><span class="token punctuation">:</span><span class="token number">23</span><span class="token punctuation">]</span> local<span class="token punctuation">.</span><span class="token constant">INFO</span><span class="token punctuation">:</span> Testing log <span class="token punctuation">{</span><span class="token string">"inputs"</span><span class="token punctuation">:</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token punctuation">,</span><span class="token string">"session"</span><span class="token punctuation">:</span><span class="token punctuation">{</span><span class="token string">"_token"</span><span class="token punctuation">:</span><span class="token string">"bZXCPViCzmaULIO6GsdLBbveC1yd3BbyH31zfT8c"</span><span class="token punctuation">,</span><span class="token string">"_previous"</span><span class="token punctuation">:</span><span class="token punctuation">{</span><span class="token string">"url"</span><span class="token punctuation">:</span>"http<span class="token punctuation">:</span><span class="token comment" spellcheck="true">//log-enhancer-demo.test"},"_flash":{"old":[],"new":[]}},"url":"/","ip":"127.0.0.1","http_method":"GET","server":"","referrer":null}</span>
</code></pre>

              <div class="postion">


                <div class="text-center ">
                  <a href="https://laravel-china.org/users/2190">
                    <img class="img-thumbnail avatar emoji" src="https://lccdn.phphub.org/uploads/avatars/2190_1439374577.png?imageView2/1/w/100/h/100" style="width:52px; height:52px;margin:5px;border: 1px solid #ddd;">
                  </a>
                  <div class="media-body padding-top-sm">
            <span class="introduction">
                  sork 翻译于 <span class="popover-with-html" data-content="2018-04-05 10:28:46" data-original-title="" title="">1周前</span>
            </span>
                  </div>
                </div>

                <div class="translation-operation-wrap">
                  <div class="index text-center">
                    <a class="section-vote btn btn-sm btn-shadow btn-default  popover-with-html" data-content="多个翻译的情况下，赞数会影响排序" data-url="https://laravel-china.org/translations/sections/2375/vote" data-original-title="" title="">
                      <i class="fa fa-thumbs-up" aria-hidden="true"></i> <span class="vote-count">0</span>
                    </a>

                    <a class="no-pjax btn btn-default btn-shadow btn-sm  popover-with-html" data-content="发现错别字？点此可帮助译者改进" href="https://laravel-china.org/sections/2375/patches/create" data-original-title="" title="">
                      <i class="fa fa-send" aria-hidden="true"></i> 改进
                    </a>

                    <a class="no-pjax btn btn-default btn-shadow btn-sm  no-pjax popover-with-html" data-content="提交你的译文" href="https://laravel-china.org/translations/sections/2375" data-original-title="" title="">
                      <i class="fa fa-paint-brush" aria-hidden="true"></i> 重译
                    </a>

                  </div>

                  <div class="text-center" style="font-size:0.9em;margin-top:12px;color:#a5a5a5;">
                    <i class="fa fa-eye" aria-hidden="true"></i> 由 <a href="https://laravel-china.org/users/1" style="color:#a5a5a5;">Summer</a> 审阅
                  </div>
                </div>
              </div>



              <div class="clearfix"></div>
            </div>
            <div class="translation-item active " style="display: block;" index="3" id="SectionIndex_3" name="SectionIndex_3">
              <p>下面是增强日志记录的配置选项（包含了默认配置选项的信息）：</p>
              <pre class="  language-php"><code class="  language-php"><span class="token delimiter">&lt;?php</span>

<span class="token keyword">return</span> <span class="token punctuation">[</span>
    <span class="token string">'log_request_details'</span> <span class="token operator">=</span><span class="token operator">&gt;</span> <span class="token boolean">true</span><span class="token punctuation">,</span>
    <span class="token string">'log_input_data'</span> <span class="token operator">=</span><span class="token operator">&gt;</span> <span class="token boolean">true</span><span class="token punctuation">,</span>
    <span class="token string">'log_request_headers'</span> <span class="token operator">=</span><span class="token operator">&gt;</span> <span class="token boolean">false</span><span class="token punctuation">,</span>
    <span class="token string">'log_session_data'</span> <span class="token operator">=</span><span class="token operator">&gt;</span> <span class="token boolean">true</span><span class="token punctuation">,</span>
    <span class="token string">'log_memory_usage'</span> <span class="token operator">=</span><span class="token operator">&gt;</span> <span class="token boolean">false</span><span class="token punctuation">,</span>
    <span class="token string">'log_git_data'</span> <span class="token operator">=</span><span class="token operator">&gt;</span> <span class="token boolean">false</span><span class="token punctuation">,</span>
    <span class="token comment" spellcheck="true">// 你可以明确不记录到日志中的用户输入信息</span>
    <span class="token string">'ignore_input_fields'</span> <span class="token operator">=</span><span class="token operator">&gt;</span> <span class="token punctuation">[</span><span class="token string">'password'</span><span class="token punctuation">,</span> <span class="token string">'confirm_password'</span><span class="token punctuation">]</span>
<span class="token punctuation">]</span><span class="token punctuation">;</span>
</code></pre>
              <p>特别注意&nbsp;<code>ignore_input_fields</code>&nbsp;选项，避免日志中记录客户的敏感信息！！！例如用户密码或信用卡账号。 </p>
              <p>你需要运行交互模式命令&nbsp;<code>artisan vendor:publish</code>&nbsp;， 以激活上面的自定义配置：</p>
              <pre class="  language-php"><code class="  language-php">php artisan vendor<span class="token punctuation">:</span>publish <span class="token operator">--</span>tag<span class="token operator">=</span>laravel<span class="token operator">-</span>log<span class="token operator">-</span>enhancer<span class="token operator">-</span>config
</code></pre>

              <div class="postion">


                <div class="text-center ">
                  <a href="https://laravel-china.org/users/2190">
                    <img class="img-thumbnail avatar emoji" src="https://lccdn.phphub.org/uploads/avatars/2190_1439374577.png?imageView2/1/w/100/h/100" style="width:52px; height:52px;margin:5px;border: 1px solid #ddd;">
                  </a>
                  <div class="media-body padding-top-sm">
            <span class="introduction">
                  sork 翻译于 <span class="popover-with-html" data-content="2018-04-05 10:39:06" data-original-title="" title="">1周前</span>
            </span>
                  </div>
                </div>

                <div class="translation-operation-wrap">
                  <div class="index text-center">
                    <a class="section-vote btn btn-sm btn-shadow btn-default  popover-with-html" data-content="多个翻译的情况下，赞数会影响排序" data-url="https://laravel-china.org/translations/sections/2376/vote" data-original-title="" title="">
                      <i class="fa fa-thumbs-up" aria-hidden="true"></i> <span class="vote-count">0</span>
                    </a>

                    <a class="no-pjax btn btn-default btn-shadow btn-sm  popover-with-html" data-content="发现错别字？点此可帮助译者改进" href="https://laravel-china.org/sections/2376/patches/create" data-original-title="" title="">
                      <i class="fa fa-send" aria-hidden="true"></i> 改进
                    </a>

                    <a class="no-pjax btn btn-default btn-shadow btn-sm  no-pjax popover-with-html" data-content="提交你的译文" href="https://laravel-china.org/translations/sections/2376" data-original-title="" title="">
                      <i class="fa fa-paint-brush" aria-hidden="true"></i> 重译
                    </a>

                  </div>

                  <div class="text-center" style="font-size:0.9em;margin-top:12px;color:#a5a5a5;">
                    <i class="fa fa-eye" aria-hidden="true"></i> 由 <a href="https://laravel-china.org/users/1" style="color:#a5a5a5;">Summer</a> 审阅
                  </div>
                </div>
              </div>



              <div class="clearfix"></div>
            </div>
            <div class="translation-item active " style="display: block;" index="4" id="SectionIndex_4" name="SectionIndex_4">
              <p>在一个基于 Laravel 5.6 的项目中，你可以使用如下命令利用 composer 来安装这个扩展包：</p>
              <pre class="  language-php"><code class="  language-php">composer <span class="token keyword">require</span> freshbitsweb<span class="token operator">/</span>laravel<span class="token operator">-</span>log<span class="token operator">-</span>enhancer
</code></pre>
              <p>由于 <a href="https://laravel-china.org/docs/laravel/5.6/packages#package-discovery"> Laravel 的包自动发现</a>&nbsp;机制，你只需要安装这个包并根据你的程序需要进行配置就可以了。</p>
              <h2>了解更多</h2>
              <p>详情请查看该包的&nbsp;<a href="https://github.com/freshbitsweb/laravel-log-enhancer">GitHub 官方仓库</a>。需要注意的是，这个扩展包依赖于 Laravel 5.6 的新日志系统，所以你只能在该版本或者更新版本中使用它。</p>
              <p>它首先出现在&nbsp;<a href="https://laravel-news.com/">Laravel News</a> 上。</p>

              <div class="postion">


                <div class="text-center ">
                  <a href="https://laravel-china.org/users/23749">
                    <img class="img-thumbnail avatar emoji" src="https://lccdn.phphub.org/uploads/avatars/23749_1522224359.jpg?imageView2/1/w/100/h/100" style="width:52px; height:52px;margin:5px;border: 1px solid #ddd;">
                  </a>
                  <div class="media-body padding-top-sm">
            <span class="introduction">
                  JiaZombie 翻译于 <span class="popover-with-html" data-content="2018-04-05 21:46:21" data-original-title="" title="">1周前</span>
            </span>
                  </div>
                </div>

                <div class="translation-operation-wrap">
                  <div class="index text-center">
                    <a class="section-vote btn btn-sm btn-shadow btn-default  popover-with-html" data-content="多个翻译的情况下，赞数会影响排序" data-url="https://laravel-china.org/translations/sections/2378/vote" data-original-title="" title="">
                      <i class="fa fa-thumbs-up" aria-hidden="true"></i> <span class="vote-count">1</span>
                    </a>

                    <a class="no-pjax btn btn-default btn-shadow btn-sm  popover-with-html" data-content="发现错别字？点此可帮助译者改进" href="https://laravel-china.org/sections/2378/patches/create" data-original-title="" title="">
                      <i class="fa fa-send" aria-hidden="true"></i> 改进
                    </a>

                    <a class="no-pjax btn btn-default btn-shadow btn-sm  no-pjax popover-with-html" data-content="提交你的译文" href="https://laravel-china.org/translations/sections/2378" data-original-title="" title="">
                      <i class="fa fa-paint-brush" aria-hidden="true"></i> 重译
                    </a>

                  </div>

                  <div class="text-center" style="font-size:0.9em;margin-top:12px;color:#a5a5a5;">
                    <i class="fa fa-eye" aria-hidden="true"></i> 由 <a href="https://laravel-china.org/users/1" style="color:#a5a5a5;">Summer</a> 审阅
                  </div>
                </div>
              </div>

              <p class="bottom-postion">

                <span class="duke-pulse view_other_version" style="left: -10px;top: 10px;"></span>

                <a class="expand-translation user-has-done-btn" data-act="view_other_version" href="javascript:;" style="margin-top: 12px;display: block;color: #a1a1a2;">
                  <i class="fa fa-history" aria-hidden="true"></i> 查看其他 2 个版本
                </a>
              </p>


              <div class="clearfix"></div>
            </div>
            <div class="translation-item  " style="display: none;" index="4" id="SectionIndex_4" name="SectionIndex_4">
              <p>你可以借由在 Laravel 5.6 专案中执行下方命令来安装此扩展包：</p>
              <pre class="  language-php"><code class="  language-php">composer <span class="token keyword">require</span> freshbitsweb<span class="token operator">/</span>laravel<span class="token operator">-</span>log<span class="token operator">-</span>enhancer
</code></pre>
              <p>得益于 Laravel 的 <a href="https://laravel-china.org/docs/laravel/5.6/packages#package-discovery">扩展包自动发现</a>，你需要做的仅只有安装本扩展包，并选择性的为你的专案进行设定。</p>
              <h2>知道更多</h2>
              <p>查看 <a href="https://github.com/freshbitsweb/laravel-log-enhancer">官方 Github 仓库</a>，注意此扩展包是基于新的 Laravel 5.6 日志，所以你必须使用最新版本的 Laravel 以使用本扩展包。</p>
              <p>这首先出现在 <a href="https://laravel-news.com/">Laravel News</a></p>
              <p>SHARE</p>

              <div class="postion">


                <div class="text-center ">
                  <a href="https://laravel-china.org/users/21335">
                    <img class="img-thumbnail avatar emoji" src="https://lccdn.phphub.org/uploads/avatars/21335_1516596331.jpg?imageView2/1/w/100/h/100" style="width:52px; height:52px;margin:5px;border: 1px solid #ddd;">
                  </a>
                  <div class="media-body padding-top-sm">
            <span class="introduction">
                  leochien 翻译于 <span class="popover-with-html" data-content="2018-04-05 21:20:22" data-original-title="" title="">1周前</span>
            </span>
                  </div>
                </div>

                <div class="translation-operation-wrap">
                  <div class="index text-center">
                    <a class="section-vote btn btn-sm btn-shadow btn-default  popover-with-html" data-content="多个翻译的情况下，赞数会影响排序" data-url="https://laravel-china.org/translations/sections/2377/vote" data-original-title="" title="">
                      <i class="fa fa-thumbs-up" aria-hidden="true"></i> <span class="vote-count">1</span>
                    </a>

                    <a class="no-pjax btn btn-default btn-shadow btn-sm  popover-with-html" data-content="发现错别字？点此可帮助译者改进" href="https://laravel-china.org/sections/2377/patches/create" data-original-title="" title="">
                      <i class="fa fa-send" aria-hidden="true"></i> 改进
                    </a>

                    <a class="no-pjax btn btn-default btn-shadow btn-sm  no-pjax popover-with-html" data-content="提交你的译文" href="https://laravel-china.org/translations/sections/2377" data-original-title="" title="">
                      <i class="fa fa-paint-brush" aria-hidden="true"></i> 重译
                    </a>

                  </div>

                  <div class="text-center" style="font-size:0.9em;margin-top:12px;color:#a5a5a5;">
                    <i class="fa fa-eye" aria-hidden="true"></i> 由 <a href="https://laravel-china.org/users/1" style="color:#a5a5a5;">Summer</a> 审阅
                  </div>
                </div>
              </div>



              <div class="clearfix"></div>
            </div>
            <div class="translation-item  " style="display: none;" index="4" id="SectionIndex_4" name="SectionIndex_4">
              <p>在一个基于 Laravel 5.6 的项目中，你可以执行以下命令利用 composer 来安装这个扩展包：</p>
              <pre class="  language-php"><code class="  language-php">composer <span class="token keyword">require</span> freshbitsweb<span class="token operator">/</span>laravel<span class="token operator">-</span>log<span class="token operator">-</span>enhancer
</code></pre>
              <p>得益于 Laravel 的自动&nbsp;<a href="https://laravel-china.org/docs/laravel/5.6/packages#package-discovery">发现扩展包机制</a>&nbsp;，你仅需要做就是安装这个扩展包，并且选择性的为你的项目程序进行配置。</p>
              <h2>了解更多</h2>
              <p>查看&nbsp;<a href="https://github.com/freshbitsweb/laravel-log-enhancer">官方 GitHub 仓库</a> 。注意此扩展包依赖于新的 Laravel 5.6 日志，因此你必须使用最新版本的 Laravel 方能使用此扩展包。</p>
              <p>这首先出现于&nbsp;<a href="https://laravel-news.com/">Laravel News</a></p>
              <p>SHARE</p>

              <div class="postion">


                <div class="text-center ">
                  <a href="https://laravel-china.org/users/18422">
                    <img class="img-thumbnail avatar emoji" src="https://lccdn.phphub.org/uploads/avatars/18422_1503113168.jpeg?imageView2/1/w/100/h/100" style="width:52px; height:52px;margin:5px;border: 1px solid #ddd;">
                  </a>
                  <div class="media-body padding-top-sm">
            <span class="introduction">
                  庞浩然Paul 翻译于 <span class="popover-with-html" data-content="2018-04-05 09:11:29" data-original-title="" title="">1周前</span>
            </span>
                  </div>
                </div>

                <div class="translation-operation-wrap">
                  <div class="index text-center">
                    <a class="section-vote btn btn-sm btn-shadow btn-default  popover-with-html" data-content="多个翻译的情况下，赞数会影响排序" data-url="https://laravel-china.org/translations/sections/2373/vote" data-original-title="" title="">
                      <i class="fa fa-thumbs-up" aria-hidden="true"></i> <span class="vote-count">0</span>
                    </a>

                    <a class="no-pjax btn btn-default btn-shadow btn-sm  popover-with-html" data-content="发现错别字？点此可帮助译者改进" href="https://laravel-china.org/sections/2373/patches/create" data-original-title="" title="">
                      <i class="fa fa-send" aria-hidden="true"></i> 改进
                    </a>

                    <a class="no-pjax btn btn-default btn-shadow btn-sm  no-pjax popover-with-html" data-content="提交你的译文" href="https://laravel-china.org/translations/sections/2373" data-original-title="" title="">
                      <i class="fa fa-paint-brush" aria-hidden="true"></i> 重译
                    </a>

                  </div>

                  <div class="text-center" style="font-size:0.9em;margin-top:12px;color:#a5a5a5;">
                    <i class="fa fa-eye" aria-hidden="true"></i> 由 <a href="https://laravel-china.org/users/1" style="color:#a5a5a5;">Summer</a> 审阅
                  </div>
                </div>
              </div>



              <div class="clearfix"></div>
            </div>
            <blockquote style="font-size: 0.9em;background: white;border: 1px dashed #dce0e4;">
              <p>原文地址：<a href="https://laravel-news.com/laravel-log-enhancer-package">https://laravel-news.com/laravel-log-enh...</a></p>
              <p>译文地址：<a href="https://laravel-china.org/topics/9422/extension-package-recommendation-laravel-log-enhancement">https://laravel-china.org/topics/9422/ex...</a></p>
            </blockquote>
            <hr>

            <blockquote style="font-size: 0.9em;">
              本文中的所有译文仅用于学习和交流目的，转载请务必注明文章译者、出处、和本文链接 <br>
              我们的翻译工作遵照 <a href="http://zh.wikipedia.org/wiki/Wikipedia:CC">CC 协议</a>，如果我们的工作有侵犯到您的权益，请及时联系我们。
            </blockquote>
            <hr>
          </div>
        </div>
      </mu-card>
    </mu-paper>
  </mu-col>
</template>

<script>
    import MuPaper from "muse-ui/src/paper/paper";
    import MuCard from "muse-ui/src/card/card";
    import MuCardHeader from "muse-ui/src/card/cardHeader";

    export default {
      components: {
        MuCardHeader,
        MuCard,
        MuPaper},
      name: "show"
    }
</script>

<style scoped>
  .article{
    margin: 16px 32px 16px 0;
  }
  .my-article-info {
    margin-left: 15px;
    color: #ccc;
    font-size: 12px;
  }
  .article-content {
    margin: 15px;
  }
</style>
