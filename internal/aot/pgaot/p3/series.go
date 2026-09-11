package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generate_series_int4_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
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
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v83 float64
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v96 float64
	_ = v96
	var v99 int32
	_ = v99
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v10 != int32(460) {
		v99 = v2
		return v99
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
		if v13 == int32(0) {
			v99 = v2
			return v99
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			if v16 != int32(15) {
				v99 = v2
				return v99
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v23 = F_estimate_expression_value(m, v19, v22)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
					v30 = F_estimate_expression_value(m, v27, v29)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
						if int32(3) <= v33 {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
							v39 = F_estimate_expression_value(m, v36, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = v39
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
								if v42 != int32(7) {
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
									if v49 != int32(7) {
										if v41 != 0 {
											v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
											if v56 != int32(7) {
												v63 = int32(0)
												if v42 != int32(7) {
													v99 = v63
												} else {
													if v49 != int32(7) {
														v99 = v63
													} else {
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
														if v68 != int32(7) {
															v99 = v63
														} else {
															v71 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
															if v71 == int32(0) {
																v99 = v63
															} else {
																v83 = base.F64_convert_i32_s(v71)
																v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
																v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
																v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
																*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
																v99 = v9
															}
														}
													}
												}
											} else {
												v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+24)))
												if v59 == int32(0) {
													v63 = int32(0)
													if v42 != int32(7) {
														v99 = v63
													} else {
														if v49 != int32(7) {
															v99 = v63
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
															if v68 != int32(7) {
																v99 = v63
															} else {
																v71 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
																if v71 == int32(0) {
																	v99 = v63
																} else {
																	v83 = base.F64_convert_i32_s(v71)
																	v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
																	v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
																	v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
																	*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
																	v99 = v9
																}
															}
														}
													}
												} else {
													v96 = float64(0)
													*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
													v99 = v9
												}
											}
										} else {
											v75 = int32(0)
											if v42 != int32(7) {
												v99 = v75
											} else {
												if v49 != int32(7) {
													v99 = v75
												} else {
													v83 = float64(1)
													v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
													v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
													*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
													v99 = v9
												}
											}
										}
									} else {
										v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+24)))
										if v52 == int32(0) {
											if v41 != 0 {
												v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
												if v56 != int32(7) {
													v63 = int32(0)
													if v42 != int32(7) {
														v99 = v63
													} else {
														if v49 != int32(7) {
															v99 = v63
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
															if v68 != int32(7) {
																v99 = v63
															} else {
																v71 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
																if v71 == int32(0) {
																	v99 = v63
																} else {
																	v83 = base.F64_convert_i32_s(v71)
																	v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
																	v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
																	v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
																	*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
																	v99 = v9
																}
															}
														}
													}
												} else {
													v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+24)))
													if v59 == int32(0) {
														v63 = int32(0)
														if v42 != int32(7) {
															v99 = v63
														} else {
															if v49 != int32(7) {
																v99 = v63
															} else {
																v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
																if v68 != int32(7) {
																	v99 = v63
																} else {
																	v71 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
																	if v71 == int32(0) {
																		v99 = v63
																	} else {
																		v83 = base.F64_convert_i32_s(v71)
																		v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
																		v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
																		v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
																		*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
																		v99 = v9
																	}
																}
															}
														}
													} else {
														v96 = float64(0)
														*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
														v99 = v9
													}
												}
											} else {
												v75 = int32(0)
												if v42 != int32(7) {
													v99 = v75
												} else {
													if v49 != int32(7) {
														v99 = v75
													} else {
														v83 = float64(1)
														v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
														v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
														v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
														*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
														v99 = v9
													}
												}
											}
										} else {
											v96 = float64(0)
											*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
											v99 = v9
										}
									}
								} else {
									v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+24)))
									if v45 == int32(0) {
										v49 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
										if v49 != int32(7) {
											if v41 != 0 {
												v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
												if v56 != int32(7) {
													v63 = int32(0)
													if v42 != int32(7) {
														v99 = v63
													} else {
														if v49 != int32(7) {
															v99 = v63
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
															if v68 != int32(7) {
																v99 = v63
															} else {
																v71 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
																if v71 == int32(0) {
																	v99 = v63
																} else {
																	v83 = base.F64_convert_i32_s(v71)
																	v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
																	v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
																	v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
																	*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
																	v99 = v9
																}
															}
														}
													}
												} else {
													v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+24)))
													if v59 == int32(0) {
														v63 = int32(0)
														if v42 != int32(7) {
															v99 = v63
														} else {
															if v49 != int32(7) {
																v99 = v63
															} else {
																v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
																if v68 != int32(7) {
																	v99 = v63
																} else {
																	v71 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
																	if v71 == int32(0) {
																		v99 = v63
																	} else {
																		v83 = base.F64_convert_i32_s(v71)
																		v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
																		v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
																		v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
																		*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
																		v99 = v9
																	}
																}
															}
														}
													} else {
														v96 = float64(0)
														*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
														v99 = v9
													}
												}
											} else {
												v75 = int32(0)
												if v42 != int32(7) {
													v99 = v75
												} else {
													if v49 != int32(7) {
														v99 = v75
													} else {
														v83 = float64(1)
														v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
														v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
														v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
														*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
														v99 = v9
													}
												}
											}
										} else {
											v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+24)))
											if v52 == int32(0) {
												if v41 != 0 {
													v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
													if v56 != int32(7) {
														v63 = int32(0)
														if v42 != int32(7) {
															v99 = v63
														} else {
															if v49 != int32(7) {
																v99 = v63
															} else {
																v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
																if v68 != int32(7) {
																	v99 = v63
																} else {
																	v71 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
																	if v71 == int32(0) {
																		v99 = v63
																	} else {
																		v83 = base.F64_convert_i32_s(v71)
																		v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
																		v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
																		v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
																		*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
																		v99 = v9
																	}
																}
															}
														}
													} else {
														v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+24)))
														if v59 == int32(0) {
															v63 = int32(0)
															if v42 != int32(7) {
																v99 = v63
															} else {
																if v49 != int32(7) {
																	v99 = v63
																} else {
																	v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
																	if v68 != int32(7) {
																		v99 = v63
																	} else {
																		v71 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
																		if v71 == int32(0) {
																			v99 = v63
																		} else {
																			v83 = base.F64_convert_i32_s(v71)
																			v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
																			v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
																			v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
																			*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
																			v99 = v9
																		}
																	}
																}
															}
														} else {
															v96 = float64(0)
															*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
															v99 = v9
														}
													}
												} else {
													v75 = int32(0)
													if v42 != int32(7) {
														v99 = v75
													} else {
														if v49 != int32(7) {
															v99 = v75
														} else {
															v83 = float64(1)
															v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
															v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
															v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
															*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
															v99 = v9
														}
													}
												}
											} else {
												v96 = float64(0)
												*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
												v99 = v9
											}
										}
									} else {
										v96 = float64(0)
										*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
										v99 = v9
									}
								}
								return v99
							}
						} else {
							v41 = int32(0)
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
							if v42 != int32(7) {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
								if v49 != int32(7) {
									if v41 != 0 {
										v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
										if v56 != int32(7) {
											v63 = int32(0)
											if v42 != int32(7) {
												v99 = v63
											} else {
												if v49 != int32(7) {
													v99 = v63
												} else {
													v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
													if v68 != int32(7) {
														v99 = v63
													} else {
														v71 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
														if v71 == int32(0) {
															v99 = v63
														} else {
															v83 = base.F64_convert_i32_s(v71)
															v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
															v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
															v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
															*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
															v99 = v9
														}
													}
												}
											}
										} else {
											v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+24)))
											if v59 == int32(0) {
												v63 = int32(0)
												if v42 != int32(7) {
													v99 = v63
												} else {
													if v49 != int32(7) {
														v99 = v63
													} else {
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
														if v68 != int32(7) {
															v99 = v63
														} else {
															v71 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
															if v71 == int32(0) {
																v99 = v63
															} else {
																v83 = base.F64_convert_i32_s(v71)
																v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
																v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
																v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
																*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
																v99 = v9
															}
														}
													}
												}
											} else {
												v96 = float64(0)
												*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
												v99 = v9
											}
										}
									} else {
										v75 = int32(0)
										if v42 != int32(7) {
											v99 = v75
										} else {
											if v49 != int32(7) {
												v99 = v75
											} else {
												v83 = float64(1)
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
												v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
												*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
												v99 = v9
											}
										}
									}
								} else {
									v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+24)))
									if v52 == int32(0) {
										if v41 != 0 {
											v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
											if v56 != int32(7) {
												v63 = int32(0)
												if v42 != int32(7) {
													v99 = v63
												} else {
													if v49 != int32(7) {
														v99 = v63
													} else {
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
														if v68 != int32(7) {
															v99 = v63
														} else {
															v71 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
															if v71 == int32(0) {
																v99 = v63
															} else {
																v83 = base.F64_convert_i32_s(v71)
																v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
																v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
																v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
																*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
																v99 = v9
															}
														}
													}
												}
											} else {
												v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+24)))
												if v59 == int32(0) {
													v63 = int32(0)
													if v42 != int32(7) {
														v99 = v63
													} else {
														if v49 != int32(7) {
															v99 = v63
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
															if v68 != int32(7) {
																v99 = v63
															} else {
																v71 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
																if v71 == int32(0) {
																	v99 = v63
																} else {
																	v83 = base.F64_convert_i32_s(v71)
																	v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
																	v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
																	v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
																	*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
																	v99 = v9
																}
															}
														}
													}
												} else {
													v96 = float64(0)
													*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
													v99 = v9
												}
											}
										} else {
											v75 = int32(0)
											if v42 != int32(7) {
												v99 = v75
											} else {
												if v49 != int32(7) {
													v99 = v75
												} else {
													v83 = float64(1)
													v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
													v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
													*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
													v99 = v9
												}
											}
										}
									} else {
										v96 = float64(0)
										*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
										v99 = v9
									}
								}
							} else {
								v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+24)))
								if v45 == int32(0) {
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
									if v49 != int32(7) {
										if v41 != 0 {
											v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
											if v56 != int32(7) {
												v63 = int32(0)
												if v42 != int32(7) {
													v99 = v63
												} else {
													if v49 != int32(7) {
														v99 = v63
													} else {
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
														if v68 != int32(7) {
															v99 = v63
														} else {
															v71 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
															if v71 == int32(0) {
																v99 = v63
															} else {
																v83 = base.F64_convert_i32_s(v71)
																v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
																v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
																v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
																*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
																v99 = v9
															}
														}
													}
												}
											} else {
												v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+24)))
												if v59 == int32(0) {
													v63 = int32(0)
													if v42 != int32(7) {
														v99 = v63
													} else {
														if v49 != int32(7) {
															v99 = v63
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
															if v68 != int32(7) {
																v99 = v63
															} else {
																v71 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
																if v71 == int32(0) {
																	v99 = v63
																} else {
																	v83 = base.F64_convert_i32_s(v71)
																	v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
																	v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
																	v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
																	*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
																	v99 = v9
																}
															}
														}
													}
												} else {
													v96 = float64(0)
													*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
													v99 = v9
												}
											}
										} else {
											v75 = int32(0)
											if v42 != int32(7) {
												v99 = v75
											} else {
												if v49 != int32(7) {
													v99 = v75
												} else {
													v83 = float64(1)
													v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
													v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
													*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
													v99 = v9
												}
											}
										}
									} else {
										v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+24)))
										if v52 == int32(0) {
											if v41 != 0 {
												v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
												if v56 != int32(7) {
													v63 = int32(0)
													if v42 != int32(7) {
														v99 = v63
													} else {
														if v49 != int32(7) {
															v99 = v63
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
															if v68 != int32(7) {
																v99 = v63
															} else {
																v71 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
																if v71 == int32(0) {
																	v99 = v63
																} else {
																	v83 = base.F64_convert_i32_s(v71)
																	v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
																	v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
																	v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
																	*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
																	v99 = v9
																}
															}
														}
													}
												} else {
													v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+24)))
													if v59 == int32(0) {
														v63 = int32(0)
														if v42 != int32(7) {
															v99 = v63
														} else {
															if v49 != int32(7) {
																v99 = v63
															} else {
																v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
																if v68 != int32(7) {
																	v99 = v63
																} else {
																	v71 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
																	if v71 == int32(0) {
																		v99 = v63
																	} else {
																		v83 = base.F64_convert_i32_s(v71)
																		v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
																		v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
																		v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
																		*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
																		v99 = v9
																	}
																}
															}
														}
													} else {
														v96 = float64(0)
														*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
														v99 = v9
													}
												}
											} else {
												v75 = int32(0)
												if v42 != int32(7) {
													v99 = v75
												} else {
													if v49 != int32(7) {
														v99 = v75
													} else {
														v83 = float64(1)
														v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
														v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
														v96 = base.F64_floor(base.F64_div(base.F64_add(v83, base.F64_sub(base.F64_convert_i32_s(v84), base.F64_convert_i32_s(v86))), v83))
														*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
														v99 = v9
													}
												}
											}
										} else {
											v96 = float64(0)
											*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
											v99 = v9
										}
									}
								} else {
									v96 = float64(0)
									*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v96
									v99 = v9
								}
							}
							return v99
						}
					}
				}
			}
		}
	}
}
func F_generate_series_int8_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
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
	var v73 int64
	_ = v73
	var v77 int32
	_ = v77
	var v84 float64
	_ = v84
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v100 float64
	_ = v100
	var v103 int32
	_ = v103
	v2 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v11 != int32(460) {
		v103 = v2
		return v103
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		if v14 == int32(0) {
			v103 = v2
			return v103
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			if v17 != int32(15) {
				v103 = v2
				return v103
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v24 = F_estimate_expression_value(m, v20, v23)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
					v31 = F_estimate_expression_value(m, v28, v30)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
						if int32(3) <= v34 {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
							v40 = F_estimate_expression_value(m, v37, v39)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v42 = v40
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
								if v43 != int32(7) {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
									if v50 != int32(7) {
										if v42 != 0 {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
											if v57 != int32(7) {
												v64 = int32(0)
												if v43 != int32(7) {
													v103 = v64
												} else {
													if v50 != int32(7) {
														v103 = v64
													} else {
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
														if v69 != int32(7) {
															v103 = v64
														} else {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
															v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
															if v73 == int64(0) {
																v103 = v64
															} else {
																v84 = base.F64_convert_i64_s(v73)
																v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
																v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
																v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
																v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
																*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
																v103 = v10
															}
														}
													}
												}
											} else {
												v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+24)))
												if v60 == int32(0) {
													v64 = int32(0)
													if v43 != int32(7) {
														v103 = v64
													} else {
														if v50 != int32(7) {
															v103 = v64
														} else {
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
															if v69 != int32(7) {
																v103 = v64
															} else {
																v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
																v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
																if v73 == int64(0) {
																	v103 = v64
																} else {
																	v84 = base.F64_convert_i64_s(v73)
																	v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
																	v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																	v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
																	v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
																	v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
																	*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
																	v103 = v10
																}
															}
														}
													}
												} else {
													v100 = float64(0)
													*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
													v103 = v10
												}
											}
										} else {
											v77 = int32(0)
											if v43 != int32(7) {
												v103 = v77
											} else {
												if v50 != int32(7) {
													v103 = v77
												} else {
													v84 = float64(1)
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
													v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
													v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
													v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
													*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
													v103 = v10
												}
											}
										}
									} else {
										v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+24)))
										if v53 == int32(0) {
											if v42 != 0 {
												v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
												if v57 != int32(7) {
													v64 = int32(0)
													if v43 != int32(7) {
														v103 = v64
													} else {
														if v50 != int32(7) {
															v103 = v64
														} else {
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
															if v69 != int32(7) {
																v103 = v64
															} else {
																v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
																v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
																if v73 == int64(0) {
																	v103 = v64
																} else {
																	v84 = base.F64_convert_i64_s(v73)
																	v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
																	v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																	v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
																	v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
																	v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
																	*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
																	v103 = v10
																}
															}
														}
													}
												} else {
													v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+24)))
													if v60 == int32(0) {
														v64 = int32(0)
														if v43 != int32(7) {
															v103 = v64
														} else {
															if v50 != int32(7) {
																v103 = v64
															} else {
																v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
																if v69 != int32(7) {
																	v103 = v64
																} else {
																	v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
																	v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
																	if v73 == int64(0) {
																		v103 = v64
																	} else {
																		v84 = base.F64_convert_i64_s(v73)
																		v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
																		v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																		v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
																		v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
																		v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
																		*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
																		v103 = v10
																	}
																}
															}
														}
													} else {
														v100 = float64(0)
														*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
														v103 = v10
													}
												}
											} else {
												v77 = int32(0)
												if v43 != int32(7) {
													v103 = v77
												} else {
													if v50 != int32(7) {
														v103 = v77
													} else {
														v84 = float64(1)
														v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
														v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
														v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
														v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
														*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
														v103 = v10
													}
												}
											}
										} else {
											v100 = float64(0)
											*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
											v103 = v10
										}
									}
								} else {
									v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+24)))
									if v46 == int32(0) {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
										if v50 != int32(7) {
											if v42 != 0 {
												v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
												if v57 != int32(7) {
													v64 = int32(0)
													if v43 != int32(7) {
														v103 = v64
													} else {
														if v50 != int32(7) {
															v103 = v64
														} else {
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
															if v69 != int32(7) {
																v103 = v64
															} else {
																v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
																v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
																if v73 == int64(0) {
																	v103 = v64
																} else {
																	v84 = base.F64_convert_i64_s(v73)
																	v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
																	v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																	v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
																	v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
																	v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
																	*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
																	v103 = v10
																}
															}
														}
													}
												} else {
													v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+24)))
													if v60 == int32(0) {
														v64 = int32(0)
														if v43 != int32(7) {
															v103 = v64
														} else {
															if v50 != int32(7) {
																v103 = v64
															} else {
																v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
																if v69 != int32(7) {
																	v103 = v64
																} else {
																	v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
																	v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
																	if v73 == int64(0) {
																		v103 = v64
																	} else {
																		v84 = base.F64_convert_i64_s(v73)
																		v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
																		v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																		v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
																		v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
																		v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
																		*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
																		v103 = v10
																	}
																}
															}
														}
													} else {
														v100 = float64(0)
														*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
														v103 = v10
													}
												}
											} else {
												v77 = int32(0)
												if v43 != int32(7) {
													v103 = v77
												} else {
													if v50 != int32(7) {
														v103 = v77
													} else {
														v84 = float64(1)
														v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
														v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
														v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
														v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
														*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
														v103 = v10
													}
												}
											}
										} else {
											v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+24)))
											if v53 == int32(0) {
												if v42 != 0 {
													v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
													if v57 != int32(7) {
														v64 = int32(0)
														if v43 != int32(7) {
															v103 = v64
														} else {
															if v50 != int32(7) {
																v103 = v64
															} else {
																v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
																if v69 != int32(7) {
																	v103 = v64
																} else {
																	v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
																	v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
																	if v73 == int64(0) {
																		v103 = v64
																	} else {
																		v84 = base.F64_convert_i64_s(v73)
																		v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
																		v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																		v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
																		v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
																		v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
																		*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
																		v103 = v10
																	}
																}
															}
														}
													} else {
														v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+24)))
														if v60 == int32(0) {
															v64 = int32(0)
															if v43 != int32(7) {
																v103 = v64
															} else {
																if v50 != int32(7) {
																	v103 = v64
																} else {
																	v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
																	if v69 != int32(7) {
																		v103 = v64
																	} else {
																		v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
																		v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
																		if v73 == int64(0) {
																			v103 = v64
																		} else {
																			v84 = base.F64_convert_i64_s(v73)
																			v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
																			v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																			v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
																			v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
																			v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
																			*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
																			v103 = v10
																		}
																	}
																}
															}
														} else {
															v100 = float64(0)
															*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
															v103 = v10
														}
													}
												} else {
													v77 = int32(0)
													if v43 != int32(7) {
														v103 = v77
													} else {
														if v50 != int32(7) {
															v103 = v77
														} else {
															v84 = float64(1)
															v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
															v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
															v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
															v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
															v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
															*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
															v103 = v10
														}
													}
												}
											} else {
												v100 = float64(0)
												*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
												v103 = v10
											}
										}
									} else {
										v100 = float64(0)
										*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
										v103 = v10
									}
								}
								return v103
							}
						} else {
							v42 = int32(0)
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
							if v43 != int32(7) {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
								if v50 != int32(7) {
									if v42 != 0 {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
										if v57 != int32(7) {
											v64 = int32(0)
											if v43 != int32(7) {
												v103 = v64
											} else {
												if v50 != int32(7) {
													v103 = v64
												} else {
													v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
													if v69 != int32(7) {
														v103 = v64
													} else {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
														v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
														if v73 == int64(0) {
															v103 = v64
														} else {
															v84 = base.F64_convert_i64_s(v73)
															v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
															v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
															v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
															v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
															v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
															*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
															v103 = v10
														}
													}
												}
											}
										} else {
											v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+24)))
											if v60 == int32(0) {
												v64 = int32(0)
												if v43 != int32(7) {
													v103 = v64
												} else {
													if v50 != int32(7) {
														v103 = v64
													} else {
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
														if v69 != int32(7) {
															v103 = v64
														} else {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
															v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
															if v73 == int64(0) {
																v103 = v64
															} else {
																v84 = base.F64_convert_i64_s(v73)
																v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
																v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
																v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
																v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
																*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
																v103 = v10
															}
														}
													}
												}
											} else {
												v100 = float64(0)
												*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
												v103 = v10
											}
										}
									} else {
										v77 = int32(0)
										if v43 != int32(7) {
											v103 = v77
										} else {
											if v50 != int32(7) {
												v103 = v77
											} else {
												v84 = float64(1)
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
												v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
												v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
												v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
												*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
												v103 = v10
											}
										}
									}
								} else {
									v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+24)))
									if v53 == int32(0) {
										if v42 != 0 {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
											if v57 != int32(7) {
												v64 = int32(0)
												if v43 != int32(7) {
													v103 = v64
												} else {
													if v50 != int32(7) {
														v103 = v64
													} else {
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
														if v69 != int32(7) {
															v103 = v64
														} else {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
															v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
															if v73 == int64(0) {
																v103 = v64
															} else {
																v84 = base.F64_convert_i64_s(v73)
																v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
																v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
																v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
																v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
																*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
																v103 = v10
															}
														}
													}
												}
											} else {
												v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+24)))
												if v60 == int32(0) {
													v64 = int32(0)
													if v43 != int32(7) {
														v103 = v64
													} else {
														if v50 != int32(7) {
															v103 = v64
														} else {
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
															if v69 != int32(7) {
																v103 = v64
															} else {
																v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
																v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
																if v73 == int64(0) {
																	v103 = v64
																} else {
																	v84 = base.F64_convert_i64_s(v73)
																	v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
																	v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																	v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
																	v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
																	v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
																	*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
																	v103 = v10
																}
															}
														}
													}
												} else {
													v100 = float64(0)
													*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
													v103 = v10
												}
											}
										} else {
											v77 = int32(0)
											if v43 != int32(7) {
												v103 = v77
											} else {
												if v50 != int32(7) {
													v103 = v77
												} else {
													v84 = float64(1)
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
													v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
													v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
													v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
													*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
													v103 = v10
												}
											}
										}
									} else {
										v100 = float64(0)
										*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
										v103 = v10
									}
								}
							} else {
								v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+24)))
								if v46 == int32(0) {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
									if v50 != int32(7) {
										if v42 != 0 {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
											if v57 != int32(7) {
												v64 = int32(0)
												if v43 != int32(7) {
													v103 = v64
												} else {
													if v50 != int32(7) {
														v103 = v64
													} else {
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
														if v69 != int32(7) {
															v103 = v64
														} else {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
															v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
															if v73 == int64(0) {
																v103 = v64
															} else {
																v84 = base.F64_convert_i64_s(v73)
																v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
																v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
																v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
																v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
																*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
																v103 = v10
															}
														}
													}
												}
											} else {
												v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+24)))
												if v60 == int32(0) {
													v64 = int32(0)
													if v43 != int32(7) {
														v103 = v64
													} else {
														if v50 != int32(7) {
															v103 = v64
														} else {
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
															if v69 != int32(7) {
																v103 = v64
															} else {
																v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
																v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
																if v73 == int64(0) {
																	v103 = v64
																} else {
																	v84 = base.F64_convert_i64_s(v73)
																	v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
																	v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																	v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
																	v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
																	v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
																	*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
																	v103 = v10
																}
															}
														}
													}
												} else {
													v100 = float64(0)
													*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
													v103 = v10
												}
											}
										} else {
											v77 = int32(0)
											if v43 != int32(7) {
												v103 = v77
											} else {
												if v50 != int32(7) {
													v103 = v77
												} else {
													v84 = float64(1)
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
													v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
													v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
													v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
													*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
													v103 = v10
												}
											}
										}
									} else {
										v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+24)))
										if v53 == int32(0) {
											if v42 != 0 {
												v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
												if v57 != int32(7) {
													v64 = int32(0)
													if v43 != int32(7) {
														v103 = v64
													} else {
														if v50 != int32(7) {
															v103 = v64
														} else {
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
															if v69 != int32(7) {
																v103 = v64
															} else {
																v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
																v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
																if v73 == int64(0) {
																	v103 = v64
																} else {
																	v84 = base.F64_convert_i64_s(v73)
																	v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
																	v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																	v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
																	v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
																	v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
																	*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
																	v103 = v10
																}
															}
														}
													}
												} else {
													v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+24)))
													if v60 == int32(0) {
														v64 = int32(0)
														if v43 != int32(7) {
															v103 = v64
														} else {
															if v50 != int32(7) {
																v103 = v64
															} else {
																v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
																if v69 != int32(7) {
																	v103 = v64
																} else {
																	v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
																	v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
																	if v73 == int64(0) {
																		v103 = v64
																	} else {
																		v84 = base.F64_convert_i64_s(v73)
																		v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
																		v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																		v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
																		v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
																		v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
																		*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
																		v103 = v10
																	}
																}
															}
														}
													} else {
														v100 = float64(0)
														*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
														v103 = v10
													}
												}
											} else {
												v77 = int32(0)
												if v43 != int32(7) {
													v103 = v77
												} else {
													if v50 != int32(7) {
														v103 = v77
													} else {
														v84 = float64(1)
														v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
														v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
														v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
														v100 = base.F64_floor(base.F64_div(base.F64_add(v84, base.F64_sub(base.F64_convert_i64_s(v87), base.F64_convert_i64_s(v90))), v84))
														*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
														v103 = v10
													}
												}
											}
										} else {
											v100 = float64(0)
											*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
											v103 = v10
										}
									}
								} else {
									v100 = float64(0)
									*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v100
									v103 = v10
								}
							}
							return v103
						}
					}
				}
			}
		}
	}
}
