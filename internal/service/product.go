package service

import (
	"context"
	"project-layout/internal/model"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// CategoryTree 获取产品分类树结构
func (s *Service) CategoryTree(ctx context.Context) ([]*model.PmsCategory, error) {
	return s.repo.CategoryAll(ctx)
}

// SkuDetail sku商品详情
func (s *Service) SkuDetail(ctx context.Context, id int) (*model.Product, error) {
	//1, sku基本信息获取
	sku, err := s.skuInfo(ctx, id)
	if err != nil {
		return nil, err
	}

	product := new(model.Product)
	// 并发执行
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		//1，获取当前spu下所有sku
		skus, err := s.repo.GetSkusBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[service.product] skus by id: %v", sku.SpuID)
		}
		product.Skus = skus
		return nil
	})
	g.Go(func() error {
		//3, sku分类下的属性分组
		attrGroups, err := s.repo.GetAttrGroupByCatID(ctx, sku.CatID)
		if err != nil {
			return errors.Wrapf(err, "[service.product] attr groups by id: %v", sku.CatID)
		}
		product.AttrGroups = attrGroups
		return nil
	})
	g.Go(func() error {
		//4, 获取spu的详情
		spu, err := s.repo.GetSpuByID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[service.product] spu info by id: %v", sku.SpuID)
		}
		if spu == nil || spu.ID == 0 {
			return ErrProductNotFound
		}
		product.Spu = spu
		return nil
	})
	g.Go(func() error {
		//5, 获取sku的销售属性组合
		skuAttrs, err := s.repo.GetSkuAttrsBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[service.product] spu skus by id: %v", sku.SpuID)
		}
		product.SkuAttrs = skuAttrs
		return nil
	})
	g.Go(func() error {
		//6, 获取spu介绍
		spuDesc, err := s.repo.GetSpuDescBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[service.product] spu desc by id: %v", sku.SpuID)
		}
		product.SpuDesc = spuDesc
		return nil
	})
	g.Go(func() error {
		//7, 获取spu的规格参数
		spuAttrs, err := s.repo.GetAttrsBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[service.product] spu attrs by id: %v", sku.SpuID)
		}
		product.SpuAttrs = spuAttrs
		return nil
	})
	if err = g.Wait(); err != nil {
		return nil, err
	}

	return product, nil
}

// GetSkuSaleAttrs 获取sku销售属性
func (s *Service) GetSkuSaleAttrs(ctx context.Context, id int) (*model.Product, error) {
	// sku基本信息获取
	sku, err := s.skuInfo(ctx, id)
	if err != nil {
		return nil, err
	}
	product := new(model.Product)

	//并发执行
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		// 获取当前spu下所有sku
		skus, err := s.repo.GetSkusBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[service.product] skus by id: %v", sku.SpuID)
		}
		product.Skus = skus
		return nil
	})
	g.Go(func() error {
		// 获取spu的详情
		spu, err := s.repo.GetSpuByID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[service.product] spu info by id: %v", sku.SpuID)
		}
		if spu == nil || spu.ID == 0 {
			return ErrProductNotFound
		}
		product.Spu = spu
		return nil
	})
	g.Go(func() error {
		//5, 获取spu的销售属性组合
		skuAttrs, err := s.repo.GetSkuAttrsBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[service.product] spu skus by id: %v", sku.SpuID)
		}
		product.SkuAttrs = skuAttrs
		return nil
	})
	if err = g.Wait(); err != nil {
		return nil, err
	}
	return product, nil
}

// GetSkuByID 获取sku信息
func (s *Service) GetSkuByID(ctx context.Context, id int) (*model.PmsSku, error) {
	return s.skuInfo(ctx, id)
}

// SpuComment 商品评价
func (s *Service) SpuComment(ctx context.Context, skuIds []int, uid, oid int, star int8, content, resources string) error {
	//sku基本信息获取
	skus, err := s.repo.GetSkusByIds(ctx, skuIds)
	if err != nil {
		return errors.Wrapf(err, "[service.product] sku info by ids: %v", skuIds)
	}
	if len(skus) == 0 {
		return ErrProductNotFound
	}
	comments := make([]*model.PmsSpuComment, 0, len(skus))
	for _, sku := range skus {
		comments = append(comments, &model.PmsSpuComment{
			SpuID:     sku.SpuID,
			SkuID:     sku.ID,
			SkuName:   sku.Name,
			MemberID:  uid,
			ReplyID:   0,
			OrderID:   oid,
			Star:      star,
			SkuAttrs:  sku.AttrValue,
			Resources: resources,
			Content:   content,
			IsRelease: 1,
		})
	}

	if err = s.repo.CreateSpuComment(ctx, comments); err != nil {
		return errors.Wrapf(err, "[service.product] batch comment")
	}
	return nil
}

// skuInfo 获取sku信息
func (s *Service) skuInfo(ctx context.Context, id int) (*model.PmsSku, error) {
	//1, sku基本信息获取
	sku, err := s.repo.GetSkuByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.product] sku info by id: %v", id)
	}
	if sku == nil || sku.ID == 0 {
		return nil, ErrProductNotFound
	}
	return sku, nil
}
