package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExplainBeginOutput(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	switch v2 - int32(1) {
	case 0:
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_appendStringInfoString(m, v5, int32(_a_F_ExplainBeginOutput_0))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v23 + int32(1)
			return
		}
	case 1:
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_appendStringInfoChar(m, v9, int32(91))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v15 = F_lcons_int(m, int32(0), v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v15
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v23 + int32(1)
				return
			}
		}
	case 2:
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v20 = F_lcons_int(m, int32(0), v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v20
			return
		}
	default:
		return
	}
}
func F_ExplainOpenWorker(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
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
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v12
	v15 = l0 << (uint(int32(4)) % 32)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v17 = v15 + v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+l0))))
	if v20 == int32(0) {
		F_initStringInfo(m, v17)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v25 + v15
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			switch v28 - int32(1) {
			case 0:
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v41 + int32(2)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				if v45 != 0 {
					F_ExplainPropertyInteger(m, int32(_a_F_ExplainOpenWorker_0), int32(0), base.I64_extend_i32_s(l0), l1)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						v53 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v51+l0))) = uint8(v53)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						if v73 == int32(0) {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
							if v77 == int32(0) {
								F_ExplainIndentText(m, l1)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
									F_appendStringInfo(m, v82, int32(_a_F_ExplainOpenWorker_1), v9)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
										m.G0 = v9 + int32(16)
										return
									}
								}
							} else {
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
								m.G0 = v9 + int32(16)
								return
							}
						} else {
							m.G0 = v9 + int32(16)
							return
						}
					}
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
					v53 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v51+l0))) = uint8(v53)
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v73 == int32(0) {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
						if v77 == int32(0) {
							F_ExplainIndentText(m, l1)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_appendStringInfo(m, v82, int32(_a_F_ExplainOpenWorker_1), v9)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
									m.G0 = v9 + int32(16)
									return
								}
							}
						} else {
							v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
							m.G0 = v9 + int32(16)
							return
						}
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				}
			case 1:
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
				v33 = F_lcons_int(m, int32(0), v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v33
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v41 + int32(2)
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v45 != 0 {
						F_ExplainPropertyInteger(m, int32(_a_F_ExplainOpenWorker_0), int32(0), base.I64_extend_i32_s(l0), l1)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							v53 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v51+l0))) = uint8(v53)
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							if v73 == int32(0) {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
								if v77 == int32(0) {
									F_ExplainIndentText(m, l1)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
										F_appendStringInfo(m, v82, int32(_a_F_ExplainOpenWorker_1), v9)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
											m.G0 = v9 + int32(16)
											return
										}
									}
								} else {
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
									m.G0 = v9 + int32(16)
									return
								}
							} else {
								m.G0 = v9 + int32(16)
								return
							}
						}
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						v53 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v51+l0))) = uint8(v53)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						if v73 == int32(0) {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
							if v77 == int32(0) {
								F_ExplainIndentText(m, l1)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
									F_appendStringInfo(m, v82, int32(_a_F_ExplainOpenWorker_1), v9)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
										m.G0 = v9 + int32(16)
										return
									}
								}
							} else {
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
								m.G0 = v9 + int32(16)
								return
							}
						} else {
							m.G0 = v9 + int32(16)
							return
						}
					}
				}
			case 2:
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
				v38 = F_lcons_int(m, int32(0), v37)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v38
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v41 + int32(2)
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v45 != 0 {
						F_ExplainPropertyInteger(m, int32(_a_F_ExplainOpenWorker_0), int32(0), base.I64_extend_i32_s(l0), l1)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							v53 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v51+l0))) = uint8(v53)
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							if v73 == int32(0) {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
								if v77 == int32(0) {
									F_ExplainIndentText(m, l1)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
										F_appendStringInfo(m, v82, int32(_a_F_ExplainOpenWorker_1), v9)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
											m.G0 = v9 + int32(16)
											return
										}
									}
								} else {
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
									m.G0 = v9 + int32(16)
									return
								}
							} else {
								m.G0 = v9 + int32(16)
								return
							}
						}
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						v53 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v51+l0))) = uint8(v53)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						if v73 == int32(0) {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
							if v77 == int32(0) {
								F_ExplainIndentText(m, l1)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
									F_appendStringInfo(m, v82, int32(_a_F_ExplainOpenWorker_1), v9)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
										m.G0 = v9 + int32(16)
										return
									}
								}
							} else {
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
								m.G0 = v9 + int32(16)
								return
							}
						} else {
							m.G0 = v9 + int32(16)
							return
						}
					}
				}
			default:
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				if v45 != 0 {
					F_ExplainPropertyInteger(m, int32(_a_F_ExplainOpenWorker_0), int32(0), base.I64_extend_i32_s(l0), l1)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						v53 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v51+l0))) = uint8(v53)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						if v73 == int32(0) {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
							if v77 == int32(0) {
								F_ExplainIndentText(m, l1)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
									F_appendStringInfo(m, v82, int32(_a_F_ExplainOpenWorker_1), v9)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
										m.G0 = v9 + int32(16)
										return
									}
								}
							} else {
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
								m.G0 = v9 + int32(16)
								return
							}
						} else {
							m.G0 = v9 + int32(16)
							return
						}
					}
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
					v53 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v51+l0))) = uint8(v53)
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v73 == int32(0) {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
						if v77 == int32(0) {
							F_ExplainIndentText(m, l1)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_appendStringInfo(m, v82, int32(_a_F_ExplainOpenWorker_1), v9)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
									m.G0 = v9 + int32(16)
									return
								}
							}
						} else {
							v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
							m.G0 = v9 + int32(16)
							return
						}
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
		v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		switch v60 - int32(1) {
		case 0:
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v68 + int32(2)
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			if v73 == int32(0) {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
				if v77 == int32(0) {
					F_ExplainIndentText(m, l1)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_appendStringInfo(m, v82, int32(_a_F_ExplainOpenWorker_1), v9)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return
						} else {
							v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
							m.G0 = v9 + int32(16)
							return
						}
					}
				} else {
					v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
					m.G0 = v9 + int32(16)
					return
				}
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		case 1, 2:
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v56+l0<<(uint(int32(2))%32))))
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			v65 = F_lcons_int(m, v63, v64)
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v65
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v68 + int32(2)
				v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				if v73 == int32(0) {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
					if v77 == int32(0) {
						F_ExplainIndentText(m, l1)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_appendStringInfo(m, v82, int32(_a_F_ExplainOpenWorker_1), v9)
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return
							} else {
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
								m.G0 = v9 + int32(16)
								return
							}
						}
					} else {
						v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
						m.G0 = v9 + int32(16)
						return
					}
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			}
		default:
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			if v73 == int32(0) {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
				if v77 == int32(0) {
					F_ExplainIndentText(m, l1)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_appendStringInfo(m, v82, int32(_a_F_ExplainOpenWorker_1), v9)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return
						} else {
							v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
							m.G0 = v9 + int32(16)
							return
						}
					}
				} else {
					v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v87 + int32(1)
					m.G0 = v9 + int32(16)
					return
				}
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		}
	}
}
func F_ExplainPreScanNode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	switch v7 - int32(335) {
	case 0:
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
		v52 = F_bms_add_members(m, v50, v51)
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int32(0)
		} else {
			v56 = v52
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v56
			v61 = F_planstate_tree_walker_impl(m, l0, int32(589), l1)
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				return v61
			}
		}
	default:
		v61 = F_planstate_tree_walker_impl(m, l0, int32(589), l1)
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return int32(0)
		} else {
			return v61
		}
	case 2:
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
		v26 = F_bms_add_member(m, v24, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v26
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+152))
			if v29 != 0 {
				v30 = F_bms_add_member(m, v26, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v30
					v33 = v30
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
					if v34 == int32(0) {
						v61 = F_planstate_tree_walker_impl(m, l0, int32(589), l1)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							return v61
						}
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+88))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
						v40 = F_bms_add_member(m, v33, v39)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							v56 = v40
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v56
							v61 = F_planstate_tree_walker_impl(m, l0, int32(589), l1)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								return v61
							}
						}
					}
				}
			} else {
				v33 = v26
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
				if v34 == int32(0) {
					v61 = F_planstate_tree_walker_impl(m, l0, int32(589), l1)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						return v61
					}
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+88))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
					v40 = F_bms_add_member(m, v33, v39)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v56 = v40
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v56
						v61 = F_planstate_tree_walker_impl(m, l0, int32(589), l1)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							return v61
						}
					}
				}
			}
		}
	case 3:
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+72))
		v44 = F_bms_add_members(m, v42, v43)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			v56 = v44
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v56
			v61 = F_planstate_tree_walker_impl(m, l0, int32(589), l1)
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				return v61
			}
		}
	case 4:
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+72))
		v48 = F_bms_add_members(m, v46, v47)
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int32(0)
		} else {
			v56 = v48
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v56
			v61 = F_planstate_tree_walker_impl(m, l0, int32(589), l1)
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				return v61
			}
		}
	case 8, 9, 10, 11, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22:
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+72))
		v12 = F_bms_add_member(m, v10, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v56 = v12
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v56
			v61 = F_planstate_tree_walker_impl(m, l0, int32(589), l1)
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				return v61
			}
		}
	case 23:
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+116))
		v18 = F_bms_add_members(m, v16, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v56 = v18
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v56
			v61 = F_planstate_tree_walker_impl(m, l0, int32(589), l1)
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				return v61
			}
		}
	case 24:
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)+100))
		v22 = F_bms_add_members(m, v20, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v56 = v22
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v56
			v61 = F_planstate_tree_walker_impl(m, l0, int32(589), l1)
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				return v61
			}
		}
	}
}
func F_ExplainPropertyUInteger(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	var v7 int32
	_ = v7
	Fn14212(m, l0, l1, l2, l3, int32(_a_F_ExplainPropertyUInteger_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_ExplainSubPlans(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 < v14 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = l1
	v23 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v12 + int32(48)
	return
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v23<<(uint(int32(2))%32))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v34 = F_bms_is_member(m, v32, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	if v34 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v40 = F_bms_add_member(m, v38, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	v76 = v18
	goto L10
L10:
	;
	v80 = v23 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v80 < v81 {
		v18 = v76
		v23 = v80
		goto L4
	} else {
		goto L25
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+48)) = v40
	v43 = F_lcons(m, v31, v18)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v45 == int32(7) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	F_ExplainNode(m, v71, v43, l2, v70, l3)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L23
	}
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v48
	v51 = F_psprintf(m, int32(_a_F_ExplainSubPlans_0), v12)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+36)))
	if v54 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v70 = v51
	goto L13
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v53
	v61 = F_psprintf(m, int32(_a_F_ExplainSubPlans_1), v12+int32(16))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v53
	v67 = F_psprintf(m, int32(_a_F_ExplainSubPlans_2), v12+int32(32))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L6
	} else {
		goto L22
	}
L21:
	;
	v70 = v61
	goto L13
L22:
	;
	v70 = v67
	goto L13
L23:
	;
	v74 = F_list_delete_first(m, v43)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v76 = v74
	goto L10
L25:
	;
	goto L5
}
