import React, {useState} from 'react';
import {Image as TaroImage, View} from '@tarojs/components';
import {XIcon} from '@/plugins/xicon';

export interface ImageProps {
    src?: string;
    fallbackSrc?: string;
    className?: string;
    mode?: 'scaleToFill' | 'aspectFit' | 'aspectFill' | 'widthFix' | 'heightFix' | 'top' | 'bottom' | 'center' | 'left' | 'right' | 'top left' | 'top right' | 'bottom left' | 'bottom right';
    style?: React.CSSProperties;
    onError?: (e: any) => void;
    [key: string]: any;
}

/**
 * 公共 Image 组件
 * - src 为空或加载失败且未提供 fallbackSrc 时，渲染纯 CSS 占位（灰底 + 文档图标）。
 *   小程序 <image> 对 SVG 支持不稳，故占位不放图片资源。
 * - 提供 fallbackSrc 时，加载失败回退到该图。
 */
const Image: React.FC<ImageProps> = ({fallbackSrc, onError, src, className, style, mode = 'aspectFill', ...rest}) => {
    const [hasError, setHasError] = useState(false);

    const handleError = (e: any) => {
        if (!hasError) {
            setHasError(true);
        }
        onError?.(e);
    };

    const resolvedSrc = hasError ? fallbackSrc : src;

    if (!resolvedSrc) {
        return (
            <View
              className={`w-full h-full flex items-center justify-center bg-pageBg ${className ?? ''}`}
              style={style}
            >
                <XIcon name='carbon:document' size={40} className='text-textWeak' />
            </View>
        );
    }

    return (
        <TaroImage
          src={resolvedSrc}
          onError={handleError}
          mode={mode}
          className={className}
          style={style}
          {...rest}
        />
    );
};

export default Image;
