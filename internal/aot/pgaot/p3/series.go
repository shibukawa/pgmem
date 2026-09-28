package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generate_series_int4_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v82 float64
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v94 float64
	_ = v94
	var v104 int64
	_ = v104
	v8 = int64(0)
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = base.I32_wrap_i64(v10)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v12 != int32(468) {
		v104 = v8
		return v104
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
		if v15 == int32(0) {
			v104 = v8
			return v104
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if v18 != int32(15) {
				v104 = v8
				return v104
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v25 = F_estimate_expression_value(m, v21, v24)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					v32 = F_estimate_expression_value(m, v29, v31)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
						if int32(3) <= v34 {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
							v40 = F_estimate_expression_value(m, v37, v39)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int64(0)
							} else {
								v42 = v40
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
								if v43 != int32(7) {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
									if v50 != int32(7) {
										if v42 != 0 {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
											if v57 != int32(7) {
												v64 = int32(7)
												if base.B2i32(v43 != v64)|base.B2i32(v50 != v64) != 0 {
													v104 = v8
												} else {
													v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
													if v69 != int32(7) {
														v104 = v8
													} else {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
														if v72 == int32(0) {
															v104 = v8
														} else {
															v82 = base.F64_convert_i32_s(v72)
															v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
															v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
															v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
															*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
															v104 = v10 & int64(4294967295)
														}
													}
												}
											} else {
												v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+32)))
												if v60 == int32(0) {
													v64 = int32(7)
													if base.B2i32(v43 != v64)|base.B2i32(v50 != v64) != 0 {
														v104 = v8
													} else {
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
														if v69 != int32(7) {
															v104 = v8
														} else {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
															if v72 == int32(0) {
																v104 = v8
															} else {
																v82 = base.F64_convert_i32_s(v72)
																v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
																v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
																v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
																*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
																v104 = v10 & int64(4294967295)
															}
														}
													}
												} else {
													v94 = float64(0)
													*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
													v104 = v10 & int64(4294967295)
												}
											}
										} else {
											if v43 != int32(7) {
												v104 = v8
											} else {
												if v50 != int32(7) {
													v104 = v8
												} else {
													v82 = float64(1)
													v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
													v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
													v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
													*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
													v104 = v10 & int64(4294967295)
												}
											}
										}
									} else {
										v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+32)))
										if v53 == int32(0) {
											if v42 != 0 {
												v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
												if v57 != int32(7) {
													v64 = int32(7)
													if base.B2i32(v43 != v64)|base.B2i32(v50 != v64) != 0 {
														v104 = v8
													} else {
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
														if v69 != int32(7) {
															v104 = v8
														} else {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
															if v72 == int32(0) {
																v104 = v8
															} else {
																v82 = base.F64_convert_i32_s(v72)
																v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
																v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
																v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
																*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
																v104 = v10 & int64(4294967295)
															}
														}
													}
												} else {
													v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+32)))
													if v60 == int32(0) {
														v64 = int32(7)
														if base.B2i32(v43 != v64)|base.B2i32(v50 != v64) != 0 {
															v104 = v8
														} else {
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
															if v69 != int32(7) {
																v104 = v8
															} else {
																v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
																if v72 == int32(0) {
																	v104 = v8
																} else {
																	v82 = base.F64_convert_i32_s(v72)
																	v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
																	v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
																	v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
																	*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
																	v104 = v10 & int64(4294967295)
																}
															}
														}
													} else {
														v94 = float64(0)
														*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
														v104 = v10 & int64(4294967295)
													}
												}
											} else {
												if v43 != int32(7) {
													v104 = v8
												} else {
													if v50 != int32(7) {
														v104 = v8
													} else {
														v82 = float64(1)
														v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
														v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
														v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
														*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
														v104 = v10 & int64(4294967295)
													}
												}
											}
										} else {
											v94 = float64(0)
											*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
											v104 = v10 & int64(4294967295)
										}
									}
								} else {
									v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+32)))
									if v46 == int32(0) {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
										if v50 != int32(7) {
											if v42 != 0 {
												v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
												if v57 != int32(7) {
													v64 = int32(7)
													if base.B2i32(v43 != v64)|base.B2i32(v50 != v64) != 0 {
														v104 = v8
													} else {
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
														if v69 != int32(7) {
															v104 = v8
														} else {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
															if v72 == int32(0) {
																v104 = v8
															} else {
																v82 = base.F64_convert_i32_s(v72)
																v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
																v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
																v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
																*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
																v104 = v10 & int64(4294967295)
															}
														}
													}
												} else {
													v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+32)))
													if v60 == int32(0) {
														v64 = int32(7)
														if base.B2i32(v43 != v64)|base.B2i32(v50 != v64) != 0 {
															v104 = v8
														} else {
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
															if v69 != int32(7) {
																v104 = v8
															} else {
																v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
																if v72 == int32(0) {
																	v104 = v8
																} else {
																	v82 = base.F64_convert_i32_s(v72)
																	v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
																	v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
																	v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
																	*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
																	v104 = v10 & int64(4294967295)
																}
															}
														}
													} else {
														v94 = float64(0)
														*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
														v104 = v10 & int64(4294967295)
													}
												}
											} else {
												if v43 != int32(7) {
													v104 = v8
												} else {
													if v50 != int32(7) {
														v104 = v8
													} else {
														v82 = float64(1)
														v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
														v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
														v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
														*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
														v104 = v10 & int64(4294967295)
													}
												}
											}
										} else {
											v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+32)))
											if v53 == int32(0) {
												if v42 != 0 {
													v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
													if v57 != int32(7) {
														v64 = int32(7)
														if base.B2i32(v43 != v64)|base.B2i32(v50 != v64) != 0 {
															v104 = v8
														} else {
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
															if v69 != int32(7) {
																v104 = v8
															} else {
																v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
																if v72 == int32(0) {
																	v104 = v8
																} else {
																	v82 = base.F64_convert_i32_s(v72)
																	v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
																	v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
																	v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
																	*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
																	v104 = v10 & int64(4294967295)
																}
															}
														}
													} else {
														v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+32)))
														if v60 == int32(0) {
															v64 = int32(7)
															if base.B2i32(v43 != v64)|base.B2i32(v50 != v64) != 0 {
																v104 = v8
															} else {
																v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
																if v69 != int32(7) {
																	v104 = v8
																} else {
																	v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
																	if v72 == int32(0) {
																		v104 = v8
																	} else {
																		v82 = base.F64_convert_i32_s(v72)
																		v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
																		v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
																		v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
																		*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
																		v104 = v10 & int64(4294967295)
																	}
																}
															}
														} else {
															v94 = float64(0)
															*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
															v104 = v10 & int64(4294967295)
														}
													}
												} else {
													if v43 != int32(7) {
														v104 = v8
													} else {
														if v50 != int32(7) {
															v104 = v8
														} else {
															v82 = float64(1)
															v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
															v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
															v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
															*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
															v104 = v10 & int64(4294967295)
														}
													}
												}
											} else {
												v94 = float64(0)
												*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
												v104 = v10 & int64(4294967295)
											}
										}
									} else {
										v94 = float64(0)
										*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
										v104 = v10 & int64(4294967295)
									}
								}
								return v104
							}
						} else {
							v42 = int32(0)
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
							if v43 != int32(7) {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
								if v50 != int32(7) {
									if v42 != 0 {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
										if v57 != int32(7) {
											v64 = int32(7)
											if base.B2i32(v43 != v64)|base.B2i32(v50 != v64) != 0 {
												v104 = v8
											} else {
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
												if v69 != int32(7) {
													v104 = v8
												} else {
													v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
													if v72 == int32(0) {
														v104 = v8
													} else {
														v82 = base.F64_convert_i32_s(v72)
														v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
														v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
														v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
														*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
														v104 = v10 & int64(4294967295)
													}
												}
											}
										} else {
											v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+32)))
											if v60 == int32(0) {
												v64 = int32(7)
												if base.B2i32(v43 != v64)|base.B2i32(v50 != v64) != 0 {
													v104 = v8
												} else {
													v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
													if v69 != int32(7) {
														v104 = v8
													} else {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
														if v72 == int32(0) {
															v104 = v8
														} else {
															v82 = base.F64_convert_i32_s(v72)
															v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
															v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
															v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
															*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
															v104 = v10 & int64(4294967295)
														}
													}
												}
											} else {
												v94 = float64(0)
												*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
												v104 = v10 & int64(4294967295)
											}
										}
									} else {
										if v43 != int32(7) {
											v104 = v8
										} else {
											if v50 != int32(7) {
												v104 = v8
											} else {
												v82 = float64(1)
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
												v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
												v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
												*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
												v104 = v10 & int64(4294967295)
											}
										}
									}
								} else {
									v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+32)))
									if v53 == int32(0) {
										if v42 != 0 {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
											if v57 != int32(7) {
												v64 = int32(7)
												if base.B2i32(v43 != v64)|base.B2i32(v50 != v64) != 0 {
													v104 = v8
												} else {
													v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
													if v69 != int32(7) {
														v104 = v8
													} else {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
														if v72 == int32(0) {
															v104 = v8
														} else {
															v82 = base.F64_convert_i32_s(v72)
															v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
															v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
															v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
															*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
															v104 = v10 & int64(4294967295)
														}
													}
												}
											} else {
												v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+32)))
												if v60 == int32(0) {
													v64 = int32(7)
													if base.B2i32(v43 != v64)|base.B2i32(v50 != v64) != 0 {
														v104 = v8
													} else {
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
														if v69 != int32(7) {
															v104 = v8
														} else {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
															if v72 == int32(0) {
																v104 = v8
															} else {
																v82 = base.F64_convert_i32_s(v72)
																v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
																v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
																v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
																*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
																v104 = v10 & int64(4294967295)
															}
														}
													}
												} else {
													v94 = float64(0)
													*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
													v104 = v10 & int64(4294967295)
												}
											}
										} else {
											if v43 != int32(7) {
												v104 = v8
											} else {
												if v50 != int32(7) {
													v104 = v8
												} else {
													v82 = float64(1)
													v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
													v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
													v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
													*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
													v104 = v10 & int64(4294967295)
												}
											}
										}
									} else {
										v94 = float64(0)
										*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
										v104 = v10 & int64(4294967295)
									}
								}
							} else {
								v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+32)))
								if v46 == int32(0) {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
									if v50 != int32(7) {
										if v42 != 0 {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
											if v57 != int32(7) {
												v64 = int32(7)
												if base.B2i32(v43 != v64)|base.B2i32(v50 != v64) != 0 {
													v104 = v8
												} else {
													v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
													if v69 != int32(7) {
														v104 = v8
													} else {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
														if v72 == int32(0) {
															v104 = v8
														} else {
															v82 = base.F64_convert_i32_s(v72)
															v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
															v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
															v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
															*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
															v104 = v10 & int64(4294967295)
														}
													}
												}
											} else {
												v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+32)))
												if v60 == int32(0) {
													v64 = int32(7)
													if base.B2i32(v43 != v64)|base.B2i32(v50 != v64) != 0 {
														v104 = v8
													} else {
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
														if v69 != int32(7) {
															v104 = v8
														} else {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
															if v72 == int32(0) {
																v104 = v8
															} else {
																v82 = base.F64_convert_i32_s(v72)
																v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
																v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
																v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
																*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
																v104 = v10 & int64(4294967295)
															}
														}
													}
												} else {
													v94 = float64(0)
													*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
													v104 = v10 & int64(4294967295)
												}
											}
										} else {
											if v43 != int32(7) {
												v104 = v8
											} else {
												if v50 != int32(7) {
													v104 = v8
												} else {
													v82 = float64(1)
													v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
													v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
													v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
													*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
													v104 = v10 & int64(4294967295)
												}
											}
										}
									} else {
										v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+32)))
										if v53 == int32(0) {
											if v42 != 0 {
												v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
												if v57 != int32(7) {
													v64 = int32(7)
													if base.B2i32(v43 != v64)|base.B2i32(v50 != v64) != 0 {
														v104 = v8
													} else {
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
														if v69 != int32(7) {
															v104 = v8
														} else {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
															if v72 == int32(0) {
																v104 = v8
															} else {
																v82 = base.F64_convert_i32_s(v72)
																v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
																v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
																v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
																*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
																v104 = v10 & int64(4294967295)
															}
														}
													}
												} else {
													v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+32)))
													if v60 == int32(0) {
														v64 = int32(7)
														if base.B2i32(v43 != v64)|base.B2i32(v50 != v64) != 0 {
															v104 = v8
														} else {
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
															if v69 != int32(7) {
																v104 = v8
															} else {
																v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
																if v72 == int32(0) {
																	v104 = v8
																} else {
																	v82 = base.F64_convert_i32_s(v72)
																	v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
																	v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
																	v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
																	*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
																	v104 = v10 & int64(4294967295)
																}
															}
														}
													} else {
														v94 = float64(0)
														*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
														v104 = v10 & int64(4294967295)
													}
												}
											} else {
												if v43 != int32(7) {
													v104 = v8
												} else {
													if v50 != int32(7) {
														v104 = v8
													} else {
														v82 = float64(1)
														v83 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
														v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
														v94 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i32_s(v83), base.F64_convert_i32_s(v85))), v82))
														*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
														v104 = v10 & int64(4294967295)
													}
												}
											}
										} else {
											v94 = float64(0)
											*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
											v104 = v10 & int64(4294967295)
										}
									}
								} else {
									v94 = float64(0)
									*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v94
									v104 = v10 & int64(4294967295)
								}
							}
							return v104
						}
					}
				}
			}
		}
	}
}
func F_generate_series_int8_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int64
	_ = v73
	var v82 float64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v95 float64
	_ = v95
	var v106 int64
	_ = v106
	v8 = int64(0)
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = base.I32_wrap_i64(v11)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v13 != int32(468) {
		v106 = v8
		return v106
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
		if v16 == int32(0) {
			v106 = v8
			return v106
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if v19 != int32(15) {
				v106 = v8
				return v106
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v26 = F_estimate_expression_value(m, v22, v25)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
					v33 = F_estimate_expression_value(m, v30, v32)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int64(0)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
						if int32(3) <= v35 {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
							v41 = F_estimate_expression_value(m, v38, v40)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int64(0)
							} else {
								v43 = v41
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
								if v44 != int32(7) {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
									if v51 != int32(7) {
										if v43 != 0 {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
											if v58 != int32(7) {
												v65 = int32(7)
												if base.B2i32(v44 != v65)|base.B2i32(v51 != v65) != 0 {
													v106 = v8
												} else {
													v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
													if v70 != int32(7) {
														v106 = v8
													} else {
														v73 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
														if v73 == int64(0) {
															v106 = v8
														} else {
															v82 = base.F64_convert_i64_s(v73)
															v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
															v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
															v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
															*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
															v106 = v11 & int64(4294967295)
														}
													}
												}
											} else {
												v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+32)))
												if v61 == int32(0) {
													v65 = int32(7)
													if base.B2i32(v44 != v65)|base.B2i32(v51 != v65) != 0 {
														v106 = v8
													} else {
														v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
														if v70 != int32(7) {
															v106 = v8
														} else {
															v73 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
															if v73 == int64(0) {
																v106 = v8
															} else {
																v82 = base.F64_convert_i64_s(v73)
																v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
																v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
																v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
																*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
																v106 = v11 & int64(4294967295)
															}
														}
													}
												} else {
													v95 = float64(0)
													*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
													v106 = v11 & int64(4294967295)
												}
											}
										} else {
											if v44 != int32(7) {
												v106 = v8
											} else {
												if v51 != int32(7) {
													v106 = v8
												} else {
													v82 = float64(1)
													v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
													v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
													v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
													*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
													v106 = v11 & int64(4294967295)
												}
											}
										}
									} else {
										v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+32)))
										if v54 == int32(0) {
											if v43 != 0 {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
												if v58 != int32(7) {
													v65 = int32(7)
													if base.B2i32(v44 != v65)|base.B2i32(v51 != v65) != 0 {
														v106 = v8
													} else {
														v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
														if v70 != int32(7) {
															v106 = v8
														} else {
															v73 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
															if v73 == int64(0) {
																v106 = v8
															} else {
																v82 = base.F64_convert_i64_s(v73)
																v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
																v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
																v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
																*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
																v106 = v11 & int64(4294967295)
															}
														}
													}
												} else {
													v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+32)))
													if v61 == int32(0) {
														v65 = int32(7)
														if base.B2i32(v44 != v65)|base.B2i32(v51 != v65) != 0 {
															v106 = v8
														} else {
															v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
															if v70 != int32(7) {
																v106 = v8
															} else {
																v73 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
																if v73 == int64(0) {
																	v106 = v8
																} else {
																	v82 = base.F64_convert_i64_s(v73)
																	v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
																	v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
																	v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
																	*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
																	v106 = v11 & int64(4294967295)
																}
															}
														}
													} else {
														v95 = float64(0)
														*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
														v106 = v11 & int64(4294967295)
													}
												}
											} else {
												if v44 != int32(7) {
													v106 = v8
												} else {
													if v51 != int32(7) {
														v106 = v8
													} else {
														v82 = float64(1)
														v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
														v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
														v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
														*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
														v106 = v11 & int64(4294967295)
													}
												}
											}
										} else {
											v95 = float64(0)
											*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
											v106 = v11 & int64(4294967295)
										}
									}
								} else {
									v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+32)))
									if v47 == int32(0) {
										v51 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
										if v51 != int32(7) {
											if v43 != 0 {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
												if v58 != int32(7) {
													v65 = int32(7)
													if base.B2i32(v44 != v65)|base.B2i32(v51 != v65) != 0 {
														v106 = v8
													} else {
														v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
														if v70 != int32(7) {
															v106 = v8
														} else {
															v73 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
															if v73 == int64(0) {
																v106 = v8
															} else {
																v82 = base.F64_convert_i64_s(v73)
																v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
																v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
																v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
																*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
																v106 = v11 & int64(4294967295)
															}
														}
													}
												} else {
													v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+32)))
													if v61 == int32(0) {
														v65 = int32(7)
														if base.B2i32(v44 != v65)|base.B2i32(v51 != v65) != 0 {
															v106 = v8
														} else {
															v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
															if v70 != int32(7) {
																v106 = v8
															} else {
																v73 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
																if v73 == int64(0) {
																	v106 = v8
																} else {
																	v82 = base.F64_convert_i64_s(v73)
																	v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
																	v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
																	v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
																	*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
																	v106 = v11 & int64(4294967295)
																}
															}
														}
													} else {
														v95 = float64(0)
														*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
														v106 = v11 & int64(4294967295)
													}
												}
											} else {
												if v44 != int32(7) {
													v106 = v8
												} else {
													if v51 != int32(7) {
														v106 = v8
													} else {
														v82 = float64(1)
														v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
														v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
														v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
														*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
														v106 = v11 & int64(4294967295)
													}
												}
											}
										} else {
											v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+32)))
											if v54 == int32(0) {
												if v43 != 0 {
													v58 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
													if v58 != int32(7) {
														v65 = int32(7)
														if base.B2i32(v44 != v65)|base.B2i32(v51 != v65) != 0 {
															v106 = v8
														} else {
															v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
															if v70 != int32(7) {
																v106 = v8
															} else {
																v73 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
																if v73 == int64(0) {
																	v106 = v8
																} else {
																	v82 = base.F64_convert_i64_s(v73)
																	v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
																	v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
																	v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
																	*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
																	v106 = v11 & int64(4294967295)
																}
															}
														}
													} else {
														v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+32)))
														if v61 == int32(0) {
															v65 = int32(7)
															if base.B2i32(v44 != v65)|base.B2i32(v51 != v65) != 0 {
																v106 = v8
															} else {
																v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
																if v70 != int32(7) {
																	v106 = v8
																} else {
																	v73 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
																	if v73 == int64(0) {
																		v106 = v8
																	} else {
																		v82 = base.F64_convert_i64_s(v73)
																		v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
																		v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
																		v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
																		*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
																		v106 = v11 & int64(4294967295)
																	}
																}
															}
														} else {
															v95 = float64(0)
															*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
															v106 = v11 & int64(4294967295)
														}
													}
												} else {
													if v44 != int32(7) {
														v106 = v8
													} else {
														if v51 != int32(7) {
															v106 = v8
														} else {
															v82 = float64(1)
															v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
															v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
															v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
															*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
															v106 = v11 & int64(4294967295)
														}
													}
												}
											} else {
												v95 = float64(0)
												*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
												v106 = v11 & int64(4294967295)
											}
										}
									} else {
										v95 = float64(0)
										*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
										v106 = v11 & int64(4294967295)
									}
								}
								return v106
							}
						} else {
							v43 = int32(0)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
							if v44 != int32(7) {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
								if v51 != int32(7) {
									if v43 != 0 {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
										if v58 != int32(7) {
											v65 = int32(7)
											if base.B2i32(v44 != v65)|base.B2i32(v51 != v65) != 0 {
												v106 = v8
											} else {
												v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
												if v70 != int32(7) {
													v106 = v8
												} else {
													v73 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
													if v73 == int64(0) {
														v106 = v8
													} else {
														v82 = base.F64_convert_i64_s(v73)
														v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
														v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
														v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
														*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
														v106 = v11 & int64(4294967295)
													}
												}
											}
										} else {
											v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+32)))
											if v61 == int32(0) {
												v65 = int32(7)
												if base.B2i32(v44 != v65)|base.B2i32(v51 != v65) != 0 {
													v106 = v8
												} else {
													v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
													if v70 != int32(7) {
														v106 = v8
													} else {
														v73 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
														if v73 == int64(0) {
															v106 = v8
														} else {
															v82 = base.F64_convert_i64_s(v73)
															v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
															v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
															v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
															*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
															v106 = v11 & int64(4294967295)
														}
													}
												}
											} else {
												v95 = float64(0)
												*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
												v106 = v11 & int64(4294967295)
											}
										}
									} else {
										if v44 != int32(7) {
											v106 = v8
										} else {
											if v51 != int32(7) {
												v106 = v8
											} else {
												v82 = float64(1)
												v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
												v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
												v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
												*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
												v106 = v11 & int64(4294967295)
											}
										}
									}
								} else {
									v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+32)))
									if v54 == int32(0) {
										if v43 != 0 {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
											if v58 != int32(7) {
												v65 = int32(7)
												if base.B2i32(v44 != v65)|base.B2i32(v51 != v65) != 0 {
													v106 = v8
												} else {
													v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
													if v70 != int32(7) {
														v106 = v8
													} else {
														v73 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
														if v73 == int64(0) {
															v106 = v8
														} else {
															v82 = base.F64_convert_i64_s(v73)
															v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
															v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
															v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
															*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
															v106 = v11 & int64(4294967295)
														}
													}
												}
											} else {
												v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+32)))
												if v61 == int32(0) {
													v65 = int32(7)
													if base.B2i32(v44 != v65)|base.B2i32(v51 != v65) != 0 {
														v106 = v8
													} else {
														v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
														if v70 != int32(7) {
															v106 = v8
														} else {
															v73 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
															if v73 == int64(0) {
																v106 = v8
															} else {
																v82 = base.F64_convert_i64_s(v73)
																v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
																v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
																v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
																*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
																v106 = v11 & int64(4294967295)
															}
														}
													}
												} else {
													v95 = float64(0)
													*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
													v106 = v11 & int64(4294967295)
												}
											}
										} else {
											if v44 != int32(7) {
												v106 = v8
											} else {
												if v51 != int32(7) {
													v106 = v8
												} else {
													v82 = float64(1)
													v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
													v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
													v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
													*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
													v106 = v11 & int64(4294967295)
												}
											}
										}
									} else {
										v95 = float64(0)
										*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
										v106 = v11 & int64(4294967295)
									}
								}
							} else {
								v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+32)))
								if v47 == int32(0) {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
									if v51 != int32(7) {
										if v43 != 0 {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
											if v58 != int32(7) {
												v65 = int32(7)
												if base.B2i32(v44 != v65)|base.B2i32(v51 != v65) != 0 {
													v106 = v8
												} else {
													v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
													if v70 != int32(7) {
														v106 = v8
													} else {
														v73 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
														if v73 == int64(0) {
															v106 = v8
														} else {
															v82 = base.F64_convert_i64_s(v73)
															v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
															v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
															v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
															*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
															v106 = v11 & int64(4294967295)
														}
													}
												}
											} else {
												v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+32)))
												if v61 == int32(0) {
													v65 = int32(7)
													if base.B2i32(v44 != v65)|base.B2i32(v51 != v65) != 0 {
														v106 = v8
													} else {
														v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
														if v70 != int32(7) {
															v106 = v8
														} else {
															v73 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
															if v73 == int64(0) {
																v106 = v8
															} else {
																v82 = base.F64_convert_i64_s(v73)
																v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
																v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
																v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
																*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
																v106 = v11 & int64(4294967295)
															}
														}
													}
												} else {
													v95 = float64(0)
													*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
													v106 = v11 & int64(4294967295)
												}
											}
										} else {
											if v44 != int32(7) {
												v106 = v8
											} else {
												if v51 != int32(7) {
													v106 = v8
												} else {
													v82 = float64(1)
													v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
													v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
													v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
													*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
													v106 = v11 & int64(4294967295)
												}
											}
										}
									} else {
										v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+32)))
										if v54 == int32(0) {
											if v43 != 0 {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
												if v58 != int32(7) {
													v65 = int32(7)
													if base.B2i32(v44 != v65)|base.B2i32(v51 != v65) != 0 {
														v106 = v8
													} else {
														v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
														if v70 != int32(7) {
															v106 = v8
														} else {
															v73 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
															if v73 == int64(0) {
																v106 = v8
															} else {
																v82 = base.F64_convert_i64_s(v73)
																v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
																v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
																v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
																*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
																v106 = v11 & int64(4294967295)
															}
														}
													}
												} else {
													v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+32)))
													if v61 == int32(0) {
														v65 = int32(7)
														if base.B2i32(v44 != v65)|base.B2i32(v51 != v65) != 0 {
															v106 = v8
														} else {
															v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
															if v70 != int32(7) {
																v106 = v8
															} else {
																v73 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
																if v73 == int64(0) {
																	v106 = v8
																} else {
																	v82 = base.F64_convert_i64_s(v73)
																	v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
																	v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
																	v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
																	*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
																	v106 = v11 & int64(4294967295)
																}
															}
														}
													} else {
														v95 = float64(0)
														*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
														v106 = v11 & int64(4294967295)
													}
												}
											} else {
												if v44 != int32(7) {
													v106 = v8
												} else {
													if v51 != int32(7) {
														v106 = v8
													} else {
														v82 = float64(1)
														v84 = *(*int64)(unsafe.Add(mBase, uint32(v33)+24))
														v86 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
														v95 = base.F64_floor(base.F64_div(base.F64_add(v82, base.F64_sub(base.F64_convert_i64_s(v84), base.F64_convert_i64_s(v86))), v82))
														*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
														v106 = v11 & int64(4294967295)
													}
												}
											}
										} else {
											v95 = float64(0)
											*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
											v106 = v11 & int64(4294967295)
										}
									}
								} else {
									v95 = float64(0)
									*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v95
									v106 = v11 & int64(4294967295)
								}
							}
							return v106
						}
					}
				}
			}
		}
	}
}
