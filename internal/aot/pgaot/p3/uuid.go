package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_uuid_decrement(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	v6 = F_palloc(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v10
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v12
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
		if v14 != 0 {
			v98 = v14
			v99 = v6 + int32(15)
			v101 = v98 - int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
			v103 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
			return v6
		} else {
			v17 = int32(255)
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)) = uint8(v17)
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)))
			if v19 != 0 {
				v98 = v19
				v99 = v6 + int32(14)
				v101 = v98 - int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
				v103 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
				return v6
			} else {
				v22 = int32(255)
				*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v22)
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)))
				if v24 != 0 {
					v98 = v24
					v99 = v6 + int32(13)
					v101 = v98 - int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
					v103 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
					return v6
				} else {
					v27 = int32(255)
					*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v27)
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
					if v29 != 0 {
						v98 = v29
						v99 = v6 + int32(12)
						v101 = v98 - int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
						v103 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
						return v6
					} else {
						v32 = int32(255)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v32)
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+11)))
						if v34 != 0 {
							v98 = v34
							v99 = v6 + int32(11)
							v101 = v98 - int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
							v103 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
							return v6
						} else {
							v37 = int32(255)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+11)) = uint8(v37)
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+10)))
							if v39 != 0 {
								v98 = v39
								v99 = v6 + int32(10)
								v101 = v98 - int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
								v103 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
								return v6
							} else {
								v42 = int32(255)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+10)) = uint8(v42)
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+9)))
								if v44 != 0 {
									v98 = v44
									v99 = v6 + int32(9)
									v101 = v98 - int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
									v103 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
									return v6
								} else {
									v47 = int32(255)
									*(*uint8)(unsafe.Add(mBase, uint32(v6)+9)) = uint8(v47)
									v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+8)))
									if v49 != 0 {
										v98 = v49
										v99 = v6 + int32(8)
										v101 = v98 - int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
										v103 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
										return v6
									} else {
										v52 = int32(255)
										*(*uint8)(unsafe.Add(mBase, uint32(v6)+8)) = uint8(v52)
										v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+7)))
										if v54 != 0 {
											v98 = v54
											v99 = v6 + int32(7)
											v101 = v98 - int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
											v103 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
											return v6
										} else {
											v57 = int32(255)
											*(*uint8)(unsafe.Add(mBase, uint32(v6)+7)) = uint8(v57)
											v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+6)))
											if v59 != 0 {
												v98 = v59
												v99 = v6 + int32(6)
												v101 = v98 - int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
												v103 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
												return v6
											} else {
												v62 = int32(255)
												*(*uint8)(unsafe.Add(mBase, uint32(v6)+6)) = uint8(v62)
												v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)))
												if v64 != 0 {
													v98 = v64
													v99 = v6 + int32(5)
													v101 = v98 - int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
													v103 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
													return v6
												} else {
													v67 = int32(255)
													*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)) = uint8(v67)
													v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)))
													if v69 != 0 {
														v98 = v69
														v99 = v6 + int32(4)
														v101 = v98 - int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
														v103 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
														return v6
													} else {
														v72 = int32(255)
														*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)) = uint8(v72)
														v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)))
														if v74 != 0 {
															v98 = v74
															v99 = v6 + int32(3)
															v101 = v98 - int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
															v103 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
															return v6
														} else {
															v77 = int32(255)
															*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)) = uint8(v77)
															v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)))
															if v79 != 0 {
																v98 = v79
																v99 = v6 + int32(2)
																v101 = v98 - int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
																v103 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
																return v6
															} else {
																v82 = int32(255)
																*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)) = uint8(v82)
																v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
																if v84 != 0 {
																	v98 = v84
																	v99 = v6 + int32(1)
																	v101 = v98 - int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
																	v103 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
																	return v6
																} else {
																	v87 = int32(255)
																	*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)) = uint8(v87)
																	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
																	if v89 != 0 {
																		v98 = v89
																		v99 = v6
																		v101 = v98 - int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
																		v103 = int32(0)
																		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
																		return v6
																	} else {
																		v90 = int32(255)
																		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v90)
																		F_pfree(m, v6)
																		mBase = m.M
																		v93 = m.ExcPending
																		if v93 != 0 {
																			return int32(0)
																		} else {
																			v94 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v94)
																			return int32(0)
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
			}
		}
	}
}
func F_uuid_fast_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v150 int64
	_ = v150
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v159 int32
	_ = v159
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v7 = int64(56)
	v9 = int64(65280)
	v11 = int64(40)
	v14 = int64(16711680)
	v16 = int64(24)
	v18 = int64(4278190080)
	v20 = int64(8)
	v41 = v6<<(uint(v7)%64) | v6&v9<<(uint(v11)%64) | (v6&v14<<(uint(v16)%64) | v6&v18<<(uint(v20)%64)) | (int64(base.Ui64(v6)>>(uint(v20)%64))&v18 | int64(base.Ui64(v6)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v6)>>(uint(v11)%64))&v9 | int64(base.Ui64(v6)>>(uint(v7)%64))))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v77 = v42<<(uint(v7)%64) | v42&v9<<(uint(v11)%64) | (v42&v14<<(uint(v16)%64) | v42&v18<<(uint(v20)%64)) | (int64(base.Ui64(v42)>>(uint(v20)%64))&v18 | int64(base.Ui64(v42)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v42)>>(uint(v11)%64))&v9 | int64(base.Ui64(v42)>>(uint(v7)%64))))
	if v41 != v77 {
		v154 = v77
		v155 = v41
		if base.Ui64(v155) < base.Ui64(v154) {
			v159 = int32(-1)
		} else {
			v159 = int32(1)
		}
		return v159
	} else {
		v79 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		v80 = int64(56)
		v82 = int64(65280)
		v84 = int64(40)
		v87 = int64(16711680)
		v89 = int64(24)
		v91 = int64(4278190080)
		v93 = int64(8)
		v114 = v79<<(uint(v80)%64) | v79&v82<<(uint(v84)%64) | (v79&v87<<(uint(v89)%64) | v79&v91<<(uint(v93)%64)) | (int64(base.Ui64(v79)>>(uint(v93)%64))&v91 | int64(base.Ui64(v79)>>(uint(v89)%64))&v87 | (int64(base.Ui64(v79)>>(uint(v84)%64))&v82 | int64(base.Ui64(v79)>>(uint(v80)%64))))
		v115 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		v150 = v115<<(uint(v80)%64) | v115&v82<<(uint(v84)%64) | (v115&v87<<(uint(v89)%64) | v115&v91<<(uint(v93)%64)) | (int64(base.Ui64(v115)>>(uint(v93)%64))&v91 | int64(base.Ui64(v115)>>(uint(v89)%64))&v87 | (int64(base.Ui64(v115)>>(uint(v84)%64))&v82 | int64(base.Ui64(v115)>>(uint(v80)%64))))
		if v114 != v150 {
			v154 = v150
			v155 = v114
			if base.Ui64(v155) < base.Ui64(v154) {
				v159 = int32(-1)
			} else {
				v159 = int32(1)
			}
			return v159
		} else {
			return int32(0)
		}
	}
}
func F_uuid_generate_v1(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v3 = m.G0
	v5 = v3 + int32(-64)
	m.G0 = v5
	F_uuid_generate_time(m, v5)
	v9 = v3 + int32(-48)
	F_uuid_unparse(m, v5, v9)
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v16 = F_DirectFunctionCall1Coll(m, int32(3376), int32(0), v9)
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			m.G0 = v5 - int32(-64)
			return v16
		}
	}
}
func F_uuid_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = int64(56)
	v9 = int64(65280)
	v11 = int64(40)
	v14 = int64(16711680)
	v16 = int64(24)
	v18 = int64(4278190080)
	v20 = int64(8)
	v41 = v6<<(uint(v7)%64) | v6&v9<<(uint(v11)%64) | (v6&v14<<(uint(v16)%64) | v6&v18<<(uint(v20)%64)) | (int64(base.Ui64(v6)>>(uint(v20)%64))&v18 | int64(base.Ui64(v6)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v6)>>(uint(v11)%64))&v9 | int64(base.Ui64(v6)>>(uint(v7)%64))))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	v78 = v43<<(uint(v7)%64) | v43&v9<<(uint(v11)%64) | (v43&v14<<(uint(v16)%64) | v43&v18<<(uint(v20)%64)) | (int64(base.Ui64(v43)>>(uint(v20)%64))&v18 | int64(base.Ui64(v43)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v43)>>(uint(v11)%64))&v9 | int64(base.Ui64(v43)>>(uint(v7)%64))))
	if v41 == v78 {
		v81 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
		v82 = int64(56)
		v84 = int64(65280)
		v86 = int64(40)
		v89 = int64(16711680)
		v91 = int64(24)
		v93 = int64(4278190080)
		v95 = int64(8)
		v116 = v81<<(uint(v82)%64) | v81&v84<<(uint(v86)%64) | (v81&v89<<(uint(v91)%64) | v81&v93<<(uint(v95)%64)) | (int64(base.Ui64(v81)>>(uint(v95)%64))&v93 | int64(base.Ui64(v81)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v81)>>(uint(v86)%64))&v84 | int64(base.Ui64(v81)>>(uint(v82)%64))))
		v117 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
		v152 = v117<<(uint(v82)%64) | v117&v84<<(uint(v86)%64) | (v117&v89<<(uint(v91)%64) | v117&v93<<(uint(v95)%64)) | (int64(base.Ui64(v117)>>(uint(v95)%64))&v93 | int64(base.Ui64(v117)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v117)>>(uint(v86)%64))&v84 | int64(base.Ui64(v117)>>(uint(v82)%64))))
		if v116 == v152 {
			v162 = int32(0)
		} else {
			v154 = v152
			v155 = v116
			if base.Ui64(v155) < base.Ui64(v154) {
				v159 = int32(-1)
			} else {
				v159 = int32(1)
			}
			v162 = v159
		}
	} else {
		v154 = v78
		v155 = v41
		if base.Ui64(v155) < base.Ui64(v154) {
			v159 = int32(-1)
		} else {
			v159 = int32(1)
		}
		v162 = v159
	}
	return base.B2i32(v162 <= int32(0))
}
func F_uuid_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v3)+8))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
	return base.B2i32(v4^v6|(v8^v9) != int64(0))
}
func F_uuid_nil(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14016(m, l0, int32(_a_F_uuid_nil_0), int32(_a_F_uuid_nil_1), int32(_a_F_uuid_nil_2), int32(_a_F_uuid_nil_3), int32(_a_F_uuid_nil_4))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
