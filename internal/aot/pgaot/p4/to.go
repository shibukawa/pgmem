package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_coerce_to_domain(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	if l1 == l3 {
		return l0
	} else {
		if l7 != 0 {
			F_hide_coercion_node(m, l0)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = F_exprTypmod(m, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					if v18 == l2 {
						v71 = l0
						v77 = F_palloc0(m, int32(28))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v77)+24)) = l6
							*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = l5
							*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v71
							*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(55)
							return v77
						}
					} else {
						if l2 < int32(0) {
							v65 = F_exprCollation(m, l0)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								v69 = F_applyRelabelType(m, l0, l1, l2, v65, int32(2), l6, int32(0))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									v71 = v69
									v77 = F_palloc0(m, int32(28))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v77)+24)) = l6
										*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = l5
										*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = int32(-1)
										*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = l3
										*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v71
										*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(55)
										return v77
									}
								}
							}
						} else {
							v23 = F_typeidType(m, l1)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return int32(0)
							} else {
								v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
								v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
								v27 = v25 + v26
								v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
								if v28 == int32(0) {
									v39 = l1
									v41 = int32(1)
								} else {
									v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
									v34 = base.B2i32(v32 == int32(_a_F_coerce_to_domain_0))
									if v32 == int32(_a_F_coerce_to_domain_0) {
										v35 = v28
									} else {
										v35 = l1
									}
									if v32 == int32(_a_F_coerce_to_domain_0) {
										v38 = int32(3)
									} else {
										v38 = int32(1)
									}
									v39 = v35
									v41 = v38
								}
								F_ReleaseCatCache(m, v23)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v45 = base.I64_extend_i32_u(v39)
									v46 = F_SearchSysCache2(m, int32(12), v45, v45)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										if v46 == int32(0) {
											v65 = F_exprCollation(m, l0)
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v69 = F_applyRelabelType(m, l0, l1, l2, v65, int32(2), l6, int32(0))
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													v71 = v69
													v77 = F_palloc0(m, int32(28))
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v77)+24)) = l6
														*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = l5
														*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = int32(-1)
														*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = l3
														*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v71
														*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(55)
														return v77
													}
												}
											}
										} else {
											v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
											v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+22)))
											v53 = *(*int32)(unsafe.Add(mBase, uint32(v50+v51)+12))
											F_ReleaseCatCache(m, v46)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return int32(0)
											} else {
												if v53 == int32(0) {
													v65 = F_exprCollation(m, l0)
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return int32(0)
													} else {
														v69 = F_applyRelabelType(m, l0, l1, l2, v65, int32(2), l6, int32(0))
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return int32(0)
														} else {
															v71 = v69
															v77 = F_palloc0(m, int32(28))
															mBase = m.M
															v78 = m.ExcPending
															if v78 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v77)+24)) = l6
																*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = l5
																*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = int32(-1)
																*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = l3
																*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v71
																*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(55)
																return v77
															}
														}
													}
												} else {
													v59 = F_build_coercion_expression(m, l0, v41, v53, l1, l2, l4, int32(2), l6)
													mBase = m.M
													v60 = m.ExcPending
													if v60 != 0 {
														return int32(0)
													} else {
														v71 = v59
														v77 = F_palloc0(m, int32(28))
														mBase = m.M
														v78 = m.ExcPending
														if v78 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v77)+24)) = l6
															*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = l5
															*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = int32(-1)
															*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = l3
															*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v71
															*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(55)
															return v77
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
			v18 = F_exprTypmod(m, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 == l2 {
					v71 = l0
					v77 = F_palloc0(m, int32(28))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v77)+24)) = l6
						*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = l5
						*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v71
						*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(55)
						return v77
					}
				} else {
					if l2 < int32(0) {
						v65 = F_exprCollation(m, l0)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							v69 = F_applyRelabelType(m, l0, l1, l2, v65, int32(2), l6, int32(0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v71 = v69
								v77 = F_palloc0(m, int32(28))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v77)+24)) = l6
									*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = l5
									*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = l3
									*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v71
									*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(55)
									return v77
								}
							}
						}
					} else {
						v23 = F_typeidType(m, l1)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
							v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
							v27 = v25 + v26
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
							if v28 == int32(0) {
								v39 = l1
								v41 = int32(1)
							} else {
								v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
								v34 = base.B2i32(v32 == int32(_a_F_coerce_to_domain_0))
								if v32 == int32(_a_F_coerce_to_domain_0) {
									v35 = v28
								} else {
									v35 = l1
								}
								if v32 == int32(_a_F_coerce_to_domain_0) {
									v38 = int32(3)
								} else {
									v38 = int32(1)
								}
								v39 = v35
								v41 = v38
							}
							F_ReleaseCatCache(m, v23)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v45 = base.I64_extend_i32_u(v39)
								v46 = F_SearchSysCache2(m, int32(12), v45, v45)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									if v46 == int32(0) {
										v65 = F_exprCollation(m, l0)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											v69 = F_applyRelabelType(m, l0, l1, l2, v65, int32(2), l6, int32(0))
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												v71 = v69
												v77 = F_palloc0(m, int32(28))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v77)+24)) = l6
													*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = l5
													*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = int32(-1)
													*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = l3
													*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v71
													*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(55)
													return v77
												}
											}
										}
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
										v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+22)))
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v50+v51)+12))
										F_ReleaseCatCache(m, v46)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int32(0)
										} else {
											if v53 == int32(0) {
												v65 = F_exprCollation(m, l0)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return int32(0)
												} else {
													v69 = F_applyRelabelType(m, l0, l1, l2, v65, int32(2), l6, int32(0))
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int32(0)
													} else {
														v71 = v69
														v77 = F_palloc0(m, int32(28))
														mBase = m.M
														v78 = m.ExcPending
														if v78 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v77)+24)) = l6
															*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = l5
															*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = int32(-1)
															*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = l3
															*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v71
															*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(55)
															return v77
														}
													}
												}
											} else {
												v59 = F_build_coercion_expression(m, l0, v41, v53, l1, l2, l4, int32(2), l6)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													v71 = v59
													v77 = F_palloc0(m, int32(28))
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v77)+24)) = l6
														*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = l5
														*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = int32(-1)
														*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = l3
														*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v71
														*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(55)
														return v77
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
func F_to_bin32(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14387(m, l0, int64(1), int64(2), int32(1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_to_date(m *base.Module, l0 int32) int64 {
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
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = int32(0)
			v28 = F_do_to_timestamp(m, v10, v15, v17, v18, v7+int32(36), v7+int32(24), v7+int32(28), v18, v18, v18)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
				if v30 <= int32(-4713) {
					if v30 != int32(-4713) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int64(0)
							} else {
								v113 = F_text_to_cstring(m, v10)
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v113
									F_errmsg(m, int32(_a_F_to_date_0), v7+int32(16))
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_to_date_1), int32(_a_F_to_date_2), int32(_a_F_to_date_3))
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int64(0)
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
								v53 = int32(_a_F_to_date_4)
							} else {
								v53 = int32(_a_F_to_date_5)
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
							v74 = base.I32_div_s((v69+v46)*int32(_a_F_to_date_6), int32(256))
							v77 = v47 + v54*int32(365) + v59 + v62 + v65 + v74 - int32(_a_F_to_date_7)
							if base.Ui32(v77) < base.Ui32(int32(2147483494)) {
								m.G0 = v7 + int32(80)
								return base.I64_extend_i32_s(v77 - int32(_a_F_to_date_8))
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return int64(0)
									} else {
										v94 = F_text_to_cstring(m, v10)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = v94
											F_errmsg(m, int32(_a_F_to_date_0), v7)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_to_date_1), int32(_a_F_to_date_9), int32(_a_F_to_date_3))
												mBase = m.M
												v104 = m.ExcPending
												if v104 != 0 {
													return int64(0)
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
							v109 = m.ExcPending
							if v109 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int64(0)
								} else {
									v113 = F_text_to_cstring(m, v10)
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v113
										F_errmsg(m, int32(_a_F_to_date_0), v7+int32(16))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_to_date_1), int32(_a_F_to_date_2), int32(_a_F_to_date_3))
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return int64(0)
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
					if v30 <= int32(_a_F_to_date_10) {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+52))
						v46 = v40
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v7)+48))
						v52 = base.B2i32(int32(2) < v46)
						if int32(2) < v46 {
							v53 = int32(_a_F_to_date_4)
						} else {
							v53 = int32(_a_F_to_date_5)
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
						v74 = base.I32_div_s((v69+v46)*int32(_a_F_to_date_6), int32(256))
						v77 = v47 + v54*int32(365) + v59 + v62 + v65 + v74 - int32(_a_F_to_date_7)
						if base.Ui32(v77) < base.Ui32(int32(2147483494)) {
							m.G0 = v7 + int32(80)
							return base.I64_extend_i32_s(v77 - int32(_a_F_to_date_8))
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int64(0)
								} else {
									v94 = F_text_to_cstring(m, v10)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v94
										F_errmsg(m, int32(_a_F_to_date_0), v7)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_to_date_1), int32(_a_F_to_date_9), int32(_a_F_to_date_3))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int64(0)
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
						if v30 != int32(_a_F_to_date_11) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int64(0)
								} else {
									v113 = F_text_to_cstring(m, v10)
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v113
										F_errmsg(m, int32(_a_F_to_date_0), v7+int32(16))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_to_date_1), int32(_a_F_to_date_2), int32(_a_F_to_date_3))
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return int64(0)
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
								v109 = m.ExcPending
								if v109 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int64(0)
									} else {
										v113 = F_text_to_cstring(m, v10)
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v113
											F_errmsg(m, int32(_a_F_to_date_0), v7+int32(16))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_to_date_1), int32(_a_F_to_date_2), int32(_a_F_to_date_3))
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int64(0)
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
									v53 = int32(_a_F_to_date_4)
								} else {
									v53 = int32(_a_F_to_date_5)
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
								v74 = base.I32_div_s((v69+v46)*int32(_a_F_to_date_6), int32(256))
								v77 = v47 + v54*int32(365) + v59 + v62 + v65 + v74 - int32(_a_F_to_date_7)
								if base.Ui32(v77) < base.Ui32(int32(2147483494)) {
									m.G0 = v7 + int32(80)
									return base.I64_extend_i32_s(v77 - int32(_a_F_to_date_8))
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return int64(0)
										} else {
											v94 = F_text_to_cstring(m, v10)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7))) = v94
												F_errmsg(m, int32(_a_F_to_date_0), v7)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_to_date_1), int32(_a_F_to_date_9), int32(_a_F_to_date_3))
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return int64(0)
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
func F_to_regoper(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14389(m, l0, int32(1691))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_to_regoperator(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14389(m, l0, int32(1692))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_to_regrole(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14389(m, l0, int32(1695))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
