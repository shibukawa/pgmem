package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_coerce_to_domain(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	if l1 == l3 {
		return l0
	} else {
		if l7 != 0 {
			F_hide_coercion_node(m, l0)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = F_exprTypmod(m, l0)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					if v17 == l2 {
						v68 = l0
						v73 = F_palloc0(m, int32(28))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = l6
							*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = l5
							*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(55)
							return v73
						}
					} else {
						if l2 < int32(0) {
							v62 = F_exprCollation(m, l0)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v66 = F_applyRelabelType(m, l0, l1, l2, v62, int32(2), l6, int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v68 = v66
									v73 = F_palloc0(m, int32(28))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = l6
										*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = l5
										*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = int32(-1)
										*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = l3
										*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v68
										*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(55)
										return v73
									}
								}
							}
						} else {
							v22 = F_typeidType(m, l1)
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return int32(0)
							} else {
								v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
								v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
								v26 = v24 + v25
								v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+92))
								if v27 == int32(0) {
									v38 = l1
									v40 = int32(1)
								} else {
									v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+88))
									v33 = base.B2i32(v31 == int32(6179))
									if v31 == int32(6179) {
										v34 = v27
									} else {
										v34 = l1
									}
									if v31 == int32(6179) {
										v37 = int32(3)
									} else {
										v37 = int32(1)
									}
									v38 = v34
									v40 = v37
								}
								F_ReleaseCatCache(m, v22)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v44 = F_SearchSysCache2(m, int32(12), v38, v38)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										if v44 == int32(0) {
											v62 = F_exprCollation(m, l0)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												v66 = F_applyRelabelType(m, l0, l1, l2, v62, int32(2), l6, int32(0))
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return int32(0)
												} else {
													v68 = v66
													v73 = F_palloc0(m, int32(28))
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = l6
														*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = l5
														*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = int32(-1)
														*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = l3
														*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v68
														*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(55)
														return v73
													}
												}
											}
										} else {
											v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
											v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+22)))
											v51 = *(*int32)(unsafe.Add(mBase, uint32(v48+v49)+12))
											F_ReleaseCatCache(m, v44)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return int32(0)
											} else {
												if v51 == int32(0) {
													v62 = F_exprCollation(m, l0)
													mBase = m.M
													v63 = m.ExcPending
													if v63 != 0 {
														return int32(0)
													} else {
														v66 = F_applyRelabelType(m, l0, l1, l2, v62, int32(2), l6, int32(0))
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return int32(0)
														} else {
															v68 = v66
															v73 = F_palloc0(m, int32(28))
															mBase = m.M
															v74 = m.ExcPending
															if v74 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = l6
																*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = l5
																*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = int32(-1)
																*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = l3
																*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v68
																*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(55)
																return v73
															}
														}
													}
												} else {
													v57 = F_build_coercion_expression(m, l0, v40, v51, l1, l2, l4, int32(2), l6)
													mBase = m.M
													v58 = m.ExcPending
													if v58 != 0 {
														return int32(0)
													} else {
														v68 = v57
														v73 = F_palloc0(m, int32(28))
														mBase = m.M
														v74 = m.ExcPending
														if v74 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = l6
															*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = l5
															*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = int32(-1)
															*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = l3
															*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v68
															*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(55)
															return v73
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v17 = F_exprTypmod(m, l0)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v17 == l2 {
					v68 = l0
					v73 = F_palloc0(m, int32(28))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = l6
						*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = l5
						*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v68
						*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(55)
						return v73
					}
				} else {
					if l2 < int32(0) {
						v62 = F_exprCollation(m, l0)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v66 = F_applyRelabelType(m, l0, l1, l2, v62, int32(2), l6, int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v68 = v66
								v73 = F_palloc0(m, int32(28))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = l6
									*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = l5
									*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = l3
									*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v68
									*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(55)
									return v73
								}
							}
						}
					} else {
						v22 = F_typeidType(m, l1)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
							v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
							v26 = v24 + v25
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+92))
							if v27 == int32(0) {
								v38 = l1
								v40 = int32(1)
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+88))
								v33 = base.B2i32(v31 == int32(6179))
								if v31 == int32(6179) {
									v34 = v27
								} else {
									v34 = l1
								}
								if v31 == int32(6179) {
									v37 = int32(3)
								} else {
									v37 = int32(1)
								}
								v38 = v34
								v40 = v37
							}
							F_ReleaseCatCache(m, v22)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v44 = F_SearchSysCache2(m, int32(12), v38, v38)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									if v44 == int32(0) {
										v62 = F_exprCollation(m, l0)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v66 = F_applyRelabelType(m, l0, l1, l2, v62, int32(2), l6, int32(0))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												v68 = v66
												v73 = F_palloc0(m, int32(28))
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = l6
													*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = l5
													*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = int32(-1)
													*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = l3
													*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v68
													*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(55)
													return v73
												}
											}
										}
									} else {
										v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
										v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+22)))
										v51 = *(*int32)(unsafe.Add(mBase, uint32(v48+v49)+12))
										F_ReleaseCatCache(m, v44)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int32(0)
										} else {
											if v51 == int32(0) {
												v62 = F_exprCollation(m, l0)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return int32(0)
												} else {
													v66 = F_applyRelabelType(m, l0, l1, l2, v62, int32(2), l6, int32(0))
													mBase = m.M
													v67 = m.ExcPending
													if v67 != 0 {
														return int32(0)
													} else {
														v68 = v66
														v73 = F_palloc0(m, int32(28))
														mBase = m.M
														v74 = m.ExcPending
														if v74 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = l6
															*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = l5
															*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = int32(-1)
															*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = l3
															*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v68
															*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(55)
															return v73
														}
													}
												}
											} else {
												v57 = F_build_coercion_expression(m, l0, v40, v51, l1, l2, l4, int32(2), l6)
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
													return int32(0)
												} else {
													v68 = v57
													v73 = F_palloc0(m, int32(28))
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = l6
														*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = l5
														*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = int32(-1)
														*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = l3
														*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v68
														*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(55)
														return v73
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_coerce_to_target_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	v9 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l2
	v23 = F_can_coerce_type(m, int32(1), v14+int32(12), v14+int32(8), l5)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v90
L2:
	;
	return int32(0)
L3:
	;
	if v23 == int32(0) {
		v90 = v9
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l1 == int32(0) {
		v55 = v9
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v58 = F_coerce_type(m, l0, v55, l2, l3, l4, l5, l6, l7)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L11
	}
L6:
	;
	v39 = l1
	goto L7
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v42 != int32(31) {
		v55 = v39
		goto L5
	} else {
		goto L9
	}
L8:
	;
	v55 = int32(0)
	goto L5
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v45 != 0 {
		v39 = v45
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	if v55 != v58 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v64 = base.B2i32(v61 != int32(7))
	goto L14
L13:
	;
	v64 = v9
	goto L14
L14:
	;
	v65 = F_coerce_type_typmod(m, v58, l3, l4, l5, l6, l7, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if l1 == v55 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v90 = v65
	goto L1
L17:
	;
	goto L18
L18:
	;
	v68 = F_type_is_collatable(m, l3)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v68 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v90 = v65
	goto L1
L21:
	;
	goto L22
L22:
	;
	v73 = F_palloc0(m, int32(16))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(31)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v80
	v90 = v73
	goto L1
}
func F_to_bin32(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	v7 = m.G0
	v8 = int32(-64)
	v9 = v7 + v8
	m.G0 = v9
	v11 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+20)))
	v13 = v9 - v8
	v14 = v13
	v15 = v11
	goto L1
L1:
	;
	v20 = int32(1)
	v21 = v14 - v20
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v15)&v20)+uint32(_consts[1353]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v27)
	if base.Ui64(v15) < base.Ui64(int64(2)) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v35 = v13 - v21
	v37 = v35 + int32(4)
	v38 = F_palloc(m, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L2
L4:
	;
	if base.Ui32(v9) < base.Ui32(v21) {
		v14 = v21
		v15 = int64(base.Ui64(v15) >> (uint(int64(1)) % 64))
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v37 << (uint(int32(2)) % 32)
	if v35 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v9 - int32(-64)
	return v38
L9:
	;
	v47 = F__emscripten_memcpy_bulkmem(m, v38+int32(4), v21, v35)
	mBase = m.M
	goto L11
L10:
	;
	goto L11
L11:
	;
	goto L8
}
func F_to_date(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = int32(0)
			v28 = F_do_to_timestamp(m, v10, v15, v17, v18, v7+int32(36), v7+int32(24), v7+int32(28), v18, v18, v18)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
				if v30 <= int32(-4713) {
					if v30 != int32(-4713) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								v112 = F_text_to_cstring(m, v10)
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v112
									F_errmsg(m, int32(727033), v7+int32(16))
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(497896), int32(4173), int32(356792))
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+52))
						if int32(10) < v35 {
							v46 = v35
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v7)+48))
							v52 = base.B2i32(int32(2) < v46)
							if int32(2) < v46 {
								v53 = int32(4800)
							} else {
								v53 = int32(4799)
							}
							v54 = v53 + v30
							v59 = base.I32_div_s(v54, int32(4))
							v62 = base.I32_div_s(v54, int32(-100))
							v65 = base.I32_div_s(v54, int32(400))
							if int32(2) < v46 {
								v69 = int32(1)
							} else {
								v69 = int32(13)
							}
							v74 = base.I32_div_s((v69+v46)*int32(7834), int32(256))
							v77 = v47 + v54*int32(365) + v59 + v62 + v65 + v74 - int32(32167)
							if base.Ui32(v77) < base.Ui32(int32(2147483494)) {
								m.G0 = v7 + int32(80)
								return v77 - int32(2451545)
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										v93 = F_text_to_cstring(m, v10)
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = v93
											F_errmsg(m, int32(727033), v7)
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(497896), int32(4182), int32(356792))
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									v112 = F_text_to_cstring(m, v10)
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v112
										F_errmsg(m, int32(727033), v7+int32(16))
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(497896), int32(4173), int32(356792))
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					if v30 <= int32(5874897) {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+52))
						v46 = v40
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v7)+48))
						v52 = base.B2i32(int32(2) < v46)
						if int32(2) < v46 {
							v53 = int32(4800)
						} else {
							v53 = int32(4799)
						}
						v54 = v53 + v30
						v59 = base.I32_div_s(v54, int32(4))
						v62 = base.I32_div_s(v54, int32(-100))
						v65 = base.I32_div_s(v54, int32(400))
						if int32(2) < v46 {
							v69 = int32(1)
						} else {
							v69 = int32(13)
						}
						v74 = base.I32_div_s((v69+v46)*int32(7834), int32(256))
						v77 = v47 + v54*int32(365) + v59 + v62 + v65 + v74 - int32(32167)
						if base.Ui32(v77) < base.Ui32(int32(2147483494)) {
							m.G0 = v7 + int32(80)
							return v77 - int32(2451545)
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									v93 = F_text_to_cstring(m, v10)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v93
										F_errmsg(m, int32(727033), v7)
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(497896), int32(4182), int32(356792))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							}
						}
					} else {
						if v30 != int32(5874898) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									v112 = F_text_to_cstring(m, v10)
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v112
										F_errmsg(m, int32(727033), v7+int32(16))
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(497896), int32(4173), int32(356792))
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							}
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+52))
							if int32(6) <= v43 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										v112 = F_text_to_cstring(m, v10)
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v112
											F_errmsg(m, int32(727033), v7+int32(16))
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(497896), int32(4173), int32(356792))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								}
							} else {
								v46 = v43
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v7)+48))
								v52 = base.B2i32(int32(2) < v46)
								if int32(2) < v46 {
									v53 = int32(4800)
								} else {
									v53 = int32(4799)
								}
								v54 = v53 + v30
								v59 = base.I32_div_s(v54, int32(4))
								v62 = base.I32_div_s(v54, int32(-100))
								v65 = base.I32_div_s(v54, int32(400))
								if int32(2) < v46 {
									v69 = int32(1)
								} else {
									v69 = int32(13)
								}
								v74 = base.I32_div_s((v69+v46)*int32(7834), int32(256))
								v77 = v47 + v54*int32(365) + v59 + v62 + v65 + v74 - int32(32167)
								if base.Ui32(v77) < base.Ui32(int32(2147483494)) {
									m.G0 = v7 + int32(80)
									return v77 - int32(2451545)
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int32(0)
										} else {
											v93 = F_text_to_cstring(m, v10)
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7))) = v93
												F_errmsg(m, int32(727033), v7)
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(497896), int32(4182), int32(356792))
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_to_regoper(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_text_to_cstring(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[1332]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v17
			v20 = *(*int64)(unsafe.Add(mBase, _consts[1333]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v20
			v26 = F_DirectInputFunctionCallSafe(m, int32(1495), v14, int32(-1), v7, v7+int32(12))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 == int32(0) {
					v30 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
					v33 = int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v33 = v32
				}
				m.G0 = v7 + int32(16)
				return v33
			}
		}
	}
}
func F_to_regoperator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_text_to_cstring(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[1332]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v17
			v20 = *(*int64)(unsafe.Add(mBase, _consts[1333]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v20
			v26 = F_DirectInputFunctionCallSafe(m, int32(1496), v14, int32(-1), v7, v7+int32(12))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 == int32(0) {
					v30 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
					v33 = int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v33 = v32
				}
				m.G0 = v7 + int32(16)
				return v33
			}
		}
	}
}
func F_to_regrole(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_text_to_cstring(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[1332]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v17
			v20 = *(*int64)(unsafe.Add(mBase, _consts[1333]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v20
			v26 = F_DirectInputFunctionCallSafe(m, int32(1499), v14, int32(-1), v7, v7+int32(12))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 == int32(0) {
					v30 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
					v33 = int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v33 = v32
				}
				m.G0 = v7 + int32(16)
				return v33
			}
		}
	}
}
