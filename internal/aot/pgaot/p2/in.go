package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_in_range_date_interval(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int64
	_ = v50
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v12 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = base.I32_wrap_i64(v13)
	v15 = int64(-9223372036854775807 - 1)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = base.I32_wrap_i64(v17)
	if v18 == int32(-2147483648) {
		v50 = v15
		if v14 == int32(-2147483648) {
			v80 = v15
			v83 = int64(0)
			v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return int64(0)
			} else {
				return v89
			}
		} else {
			if v14 == int32(2147483647) {
				v80 = int64(9223372036854775807)
				v83 = int64(0)
				v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int64(0)
				} else {
					return v89
				}
			} else {
				if int32(106751983) <= v14 {
					v58 = int64(9223372036854775807)
					v60 = F_errsave_start(m, int32(0))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int64(0)
					} else {
						if v60 == int32(0) {
							v80 = v58
							v83 = int64(0)
							v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int64(0)
							} else {
								return v89
							}
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_in_range_date_interval_0), int32(0))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int64(0)
								} else {
									F_errsave_finish(m, int32(0), int32(_a_F_in_range_date_interval_1), int32(646), int32(_a_F_in_range_date_interval_2))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int64(0)
									} else {
										v80 = v58
										v83 = int64(0)
										v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int64(0)
										} else {
											return v89
										}
									}
								}
							}
						}
					}
				} else {
					v80 = base.I64_extend32_s(v13) * int64(86400000000)
					v83 = int64(0)
					v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int64(0)
					} else {
						return v89
					}
				}
			}
		}
	} else {
		if v18 == int32(2147483647) {
			v50 = int64(9223372036854775807)
			if v14 == int32(-2147483648) {
				v80 = v15
				v83 = int64(0)
				v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int64(0)
				} else {
					return v89
				}
			} else {
				if v14 == int32(2147483647) {
					v80 = int64(9223372036854775807)
					v83 = int64(0)
					v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int64(0)
					} else {
						return v89
					}
				} else {
					if int32(106751983) <= v14 {
						v58 = int64(9223372036854775807)
						v60 = F_errsave_start(m, int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int64(0)
						} else {
							if v60 == int32(0) {
								v80 = v58
								v83 = int64(0)
								v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return int64(0)
								} else {
									return v89
								}
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_in_range_date_interval_0), int32(0))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int64(0)
									} else {
										F_errsave_finish(m, int32(0), int32(_a_F_in_range_date_interval_1), int32(646), int32(_a_F_in_range_date_interval_2))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int64(0)
										} else {
											v80 = v58
											v83 = int64(0)
											v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int64(0)
											} else {
												return v89
											}
										}
									}
								}
							}
						}
					} else {
						v80 = base.I64_extend32_s(v13) * int64(86400000000)
						v83 = int64(0)
						v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int64(0)
						} else {
							return v89
						}
					}
				}
			}
		} else {
			if int32(106751983) <= v18 {
				v26 = int64(9223372036854775807)
				v28 = F_errsave_start(m, int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					if v28 == int32(0) {
						v50 = v26
						if v14 == int32(-2147483648) {
							v80 = v15
							v83 = int64(0)
							v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int64(0)
							} else {
								return v89
							}
						} else {
							if v14 == int32(2147483647) {
								v80 = int64(9223372036854775807)
								v83 = int64(0)
								v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return int64(0)
								} else {
									return v89
								}
							} else {
								if int32(106751983) <= v14 {
									v58 = int64(9223372036854775807)
									v60 = F_errsave_start(m, int32(0))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int64(0)
									} else {
										if v60 == int32(0) {
											v80 = v58
											v83 = int64(0)
											v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int64(0)
											} else {
												return v89
											}
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_in_range_date_interval_0), int32(0))
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int64(0)
												} else {
													F_errsave_finish(m, int32(0), int32(_a_F_in_range_date_interval_1), int32(646), int32(_a_F_in_range_date_interval_2))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int64(0)
													} else {
														v80 = v58
														v83 = int64(0)
														v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
														mBase = m.M
														v90 = m.ExcPending
														if v90 != 0 {
															return int64(0)
														} else {
															return v89
														}
													}
												}
											}
										}
									}
								} else {
									v80 = base.I64_extend32_s(v13) * int64(86400000000)
									v83 = int64(0)
									v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return int64(0)
									} else {
										return v89
									}
								}
							}
						}
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_in_range_date_interval_0), int32(0))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, int32(0), int32(_a_F_in_range_date_interval_1), int32(646), int32(_a_F_in_range_date_interval_2))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int64(0)
								} else {
									v50 = v26
									if v14 == int32(-2147483648) {
										v80 = v15
										v83 = int64(0)
										v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int64(0)
										} else {
											return v89
										}
									} else {
										if v14 == int32(2147483647) {
											v80 = int64(9223372036854775807)
											v83 = int64(0)
											v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int64(0)
											} else {
												return v89
											}
										} else {
											if int32(106751983) <= v14 {
												v58 = int64(9223372036854775807)
												v60 = F_errsave_start(m, int32(0))
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return int64(0)
												} else {
													if v60 == int32(0) {
														v80 = v58
														v83 = int64(0)
														v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
														mBase = m.M
														v90 = m.ExcPending
														if v90 != 0 {
															return int64(0)
														} else {
															return v89
														}
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v66 = m.ExcPending
														if v66 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_in_range_date_interval_0), int32(0))
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return int64(0)
															} else {
																F_errsave_finish(m, int32(0), int32(_a_F_in_range_date_interval_1), int32(646), int32(_a_F_in_range_date_interval_2))
																mBase = m.M
																v76 = m.ExcPending
																if v76 != 0 {
																	return int64(0)
																} else {
																	v80 = v58
																	v83 = int64(0)
																	v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
																	mBase = m.M
																	v90 = m.ExcPending
																	if v90 != 0 {
																		return int64(0)
																	} else {
																		return v89
																	}
																}
															}
														}
													}
												}
											} else {
												v80 = base.I64_extend32_s(v13) * int64(86400000000)
												v83 = int64(0)
												v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return int64(0)
												} else {
													return v89
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
				v50 = base.I64_extend32_s(v17) * int64(86400000000)
				if v14 == int32(-2147483648) {
					v80 = v15
					v83 = int64(0)
					v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int64(0)
					} else {
						return v89
					}
				} else {
					if v14 == int32(2147483647) {
						v80 = int64(9223372036854775807)
						v83 = int64(0)
						v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int64(0)
						} else {
							return v89
						}
					} else {
						if int32(106751983) <= v14 {
							v58 = int64(9223372036854775807)
							v60 = F_errsave_start(m, int32(0))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int64(0)
							} else {
								if v60 == int32(0) {
									v80 = v58
									v83 = int64(0)
									v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return int64(0)
									} else {
										return v89
									}
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_in_range_date_interval_0), int32(0))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int64(0)
										} else {
											F_errsave_finish(m, int32(0), int32(_a_F_in_range_date_interval_1), int32(646), int32(_a_F_in_range_date_interval_2))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int64(0)
											} else {
												v80 = v58
												v83 = int64(0)
												v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return int64(0)
												} else {
													return v89
												}
											}
										}
									}
								}
							}
						} else {
							v80 = base.I64_extend32_s(v13) * int64(86400000000)
							v83 = int64(0)
							v89 = F_DirectFunctionCall5Coll(m, int32(1394), int32(0), v50, v80, v12, base.I64_extend_i32_u(base.B2i32(v11 != v83)), base.I64_extend_i32_u(base.B2i32(v10 != v83)))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int64(0)
							} else {
								return v89
							}
						}
					}
				}
			}
		}
	}
}
